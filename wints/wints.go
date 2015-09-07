package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/httpd"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/sqlstore"
	_ "github.com/lib/pq"
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

func startFeeder(cfg config.Feeder, store *sqlstore.Store) {

	r := feeder.NewHTTPConventionReader(cfg.URL, cfg.Login, cfg.Password)
	r.Encoding = cfg.Encoding

	f := feeder.NewCsvConventions(r, cfg.Promotions)
	go f.Import(store)
	ticker := time.NewTicker(cfg.Frequency.Duration)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				f.Import(store)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
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

	//the feeder
	startFeeder(cfg.Feeder, store)
	log.Println("Convention feeder started")
	httpd := httpd.NewHTTPd(store, cfg.HTTPd, cfg.Internships)
	log.Fatalf("%s\n", httpd.Listen())
}
