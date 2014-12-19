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

func setup() (config.Config, mail.Mailer, *datastore.Service) {
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

	if *reset {
		err = datastore.Reset(DB)
		if err != nil {
			log.Fatalln("Unable to reset the root account: " + err.Error())
		}
		log.Println("Root account reset. Don't forgot to delete it once logged")
	}
	return cfg, mailer, ds
}
func main() {
	_, _, ds := setup()
	defer ds.Db.Close()

	/*	for _, v := range os.Args {
		if v == "-s" {
			delay, err := time.ParseDuration(cfg.Puller.Period)
			if err != nil {
				log.Fatalln("Unable to parse refresh duration: " + err.Error())
			}
			backend.DaemonConventionsPuller(DB, cfg.Puller.Url,
				cfg.Puller.Login,
				cfg.Puller.Password,
				delay)
		}
	}*/

}
