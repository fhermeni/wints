package main

import (
	"database/sql"
	"errors"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/httpd"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/spy"
	"github.com/fhermeni/wints/sqlstore"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
)

var cfg config.Config

func inviteRoot(store *sqlstore.Store, em string) error {
	p := schema.Person{Firstname: "root", Lastname: "root", Email: em, Tel: "n/a"}
	if err := store.NewUser(p, schema.ROOT); err != nil {
		log.Println(err.Error())
		return err
	}
	tok, err := store.ResetPassword(p.Email, cfg.HTTPd.Rest.RenewalRequestLifetime.Duration)
	if err != nil {
		log.Println("Unable to prepare the password reset:%s\n", err.Error())
		return err
	}
	//{{.WWW}}/resetPassword?token={{.Token}}
	log.Println(cfg.HTTPd.WWW + "/password.html?resetToken=" + string(tok))
	return nil
}

func newConventionReader(cfg config.Feeder) feeder.ConventionReader {
	reader := feeder.NewHTTPConventionReader(cfg.URL, cfg.Login, cfg.Password)
	reader.Encoding = cfg.Encoding
	return reader
}

func newFeeder(cfg config.Feeder) feeder.Conventions {
	r := feeder.NewHTTPConventionReader(cfg.URL, cfg.Login, cfg.Password)
	r.Encoding = cfg.Encoding

	f := feeder.NewCsvConventions(r, cfg.Promotions)
	return f
}

func runSpies(st *sqlstore.Store, cfg config.Config) error {

	if len(cfg.Spies) == 0 {
		log.Println("No spies to schedule")
		return nil
	}
	cr := cron.New()
	for _, s := range cfg.Spies {
		switch strings.ToLower(s.Kind) {
		case "tutor":
			reviews := make(map[string]time.Duration)
			for k, r := range cfg.Internships.Reports {
				reviews[k] = r.Review.Duration
			}
			cr.AddJob(s.Cron, spy.NewTutor(st, reviews))
		default:
			return errors.New("Unsupported spy '" + s.Kind + "'")
		}
		log.Println("Spy '" + s.Kind + "' scheduled")
	}
	cr.Start()
	return nil
}

func main() {

	makeRoot := flag.String("invite-root", "", "Invite a root user")
	flag.Parse()

	if _, err := toml.DecodeFile("wints_test.conf", &cfg); err != nil {
		log.Fatalln(err.Error())
	}

	//the store
	DB, err := sql.Open("postgres", cfg.Db.ConnectionString)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}
	store, _ := sqlstore.NewStore(DB, cfg.Internships)

	if len(*makeRoot) > 0 {
		inviteRoot(store, *makeRoot)
		os.Exit(0)
	}

	//the spies
	err = runSpies(store, cfg)
	if err != nil {
		log.Fatalln("Unable to launch the spies: " + err.Error())
	}
	//the feeder
	conventions := newFeeder(cfg.Feeder)
	httpd := httpd.NewHTTPd(store, conventions, cfg.HTTPd, cfg.Internships)
	log.Fatalf("%s\n", httpd.Listen())
}
