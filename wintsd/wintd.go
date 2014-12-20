//launch wints daemon
package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/datastore"
	"github.com/fhermeni/wints/mail"

	_ "github.com/lib/pq"
)

const (
	DEFAULT_LOGIN    = "root@localhost.com"
	DEFAULT_PASSWORD = "wints"
)

func setup() (config.Config, mail.Mailer, *datastore.Service, bool) {
	cfgPath := flag.String("conf", "./wints.conf", "daemon configuration file")
	reset := flag.Bool("reset", false, "Reset the root account")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalln("Error while parsing " + *cfgPath + ": " + err.Error())
	}
	log.Println("Parsing " + *cfgPath + ": OK")

	mailer := mail.NewMailer(cfg.Mailer.Server, cfg.Mailer.Login, cfg.Mailer.Password, cfg.Mailer.Sender)
	if err != nil {
		log.Fatalln("Unable to setup the mailer: " + err.Error())
	}

	DB, err := sql.Open("postgres", cfg.Db.Url)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}
	ds, err := datastore.NewService(DB)
	if err != nil {
		log.Fatalln("Unable to test the connection: " + err.Error())
	}
	log.Println("Database connection: OK")

	return cfg, mailer, ds, *reset
}
func main() {
	_, _, ds, reset := setup()
	defer ds.Db.Close()

	if reset {
		err := datastore.Reset(ds.Db, DEFAULT_LOGIN, DEFAULT_PASSWORD)
		if err != nil {
			log.Fatalln("Unable to reset the root account: " + err.Error())
		}
		log.Println("Root account reset. Don't forgot to delete it once logged")
	}
}
