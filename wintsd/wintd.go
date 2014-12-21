//launch wints daemon
package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/datastore"
	"github.com/fhermeni/wints/mail"

	_ "github.com/lib/pq"
)

func confirm() bool {
	os.Stdout.WriteString("Confirm (y/n) ?")
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	ret := string(b) == "y"
	if !ret {
		os.Stdout.WriteString("Aborded")
	}
	return ret
}

/*
 --install -> check there is no user table. Ok -> install.sql
 --force-install -> clean.sql & install.sql
 --root-account
*/
func setup() (config.Config, mail.Mailer, *datastore.Service, bool, bool, bool) {
	cfgPath := flag.String("conf", "./wints.conf", "daemon configuration file")
	reset := flag.Bool("reset", false, "Reset the root account")
	install := flag.Bool("install", false, "/!\\ Create database tables")
	clean := flag.Bool("clean", false, "/!\\ Drop tables")
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

	return cfg, mailer, ds, *reset, *clean, *install
}

func main() {
	_, _, ds, reset, clean, install := setup()
	defer ds.Db.Close()

	if (install || clean) && confirm() {
		if clean {
			err := ds.Clean()
			if err != nil {
				log.Fatalln("Unable to clean the tables: " + err.Error())
			}
			log.Println("Table removed")
		}
		if install {
			err := ds.Install()
			if err != nil {
				log.Fatalln("Unable to create the tables: " + err.Error())
			}
			log.Println("Table created")
		}
		os.Exit(0)
	}

	if reset {
		err := ds.ResetRootAccount()
		if err != nil {
			log.Fatalln("Unable to reset the root account: " + err.Error())
		}
		log.Println("Root account reset. Don't forgot to delete it once logged")
	}
}
