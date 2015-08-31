package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fhermeni/wints/config"
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
	store, _ := sqlstore.NewStore(DB)

	if len(*makeRoot) > 0 {
		inviteRoot(store, *makeRoot)
		os.Exit(0)
	}

	httpd := httpd.NewHTTPd(store, cfg.HTTPd, cfg.Internships)
	log.Fatalf("%s\n", httpd.Listen())
}
