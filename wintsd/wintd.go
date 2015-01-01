//launch wints daemon
package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/datastore"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/handler"
	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/mail"

	_ "github.com/lib/pq"
)

func confirm(msg string) bool {
	os.Stdout.WriteString(msg + " (y/n) ?")
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	ret := string(b) == "y"
	if !ret {
		os.Stdout.WriteString("Aborded")
	}
	return ret
}

/*
 --test-db
 --invitation XXX
 --install
 --test-mailer
 --test-feeder
*/
func setup() (config.Config, mail.SMTP, *datastore.Service, bool, bool) {
	cfgPath := flag.String("conf", "./wints.conf", "daemon configuration file")
	reset := flag.Bool("reset", false, "Reset the root account")
	install := flag.Bool("install", false, "/!\\ Create database tables")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalln("Error while parsing " + *cfgPath + ": " + err.Error())
	}
	log.Println("Parsing " + *cfgPath + ": OK")

	mailer := mail.NewSMTP(cfg.Mailer.Server, cfg.Mailer.Login, cfg.Mailer.Password, cfg.Mailer.Sender)
	if err != nil {
		log.Fatalln("Unable to setup the mailer: " + err.Error())
	}

	DB, err := sql.Open("postgres", cfg.DB.URL)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}
	ds, err := datastore.NewService(DB)
	if err != nil {
		log.Fatalln("Unable to test the connection: " + err.Error())
	}
	log.Println("Database connection: OK")

	return cfg, mailer, ds, *reset, *install
}

func startFeederDaemon(f feeder.Feeder, srv internship.Service, frequency time.Duration) {
	log.Println("Scanning conventions")
	go f.InjectConventions(srv)
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Scanning conventions")
				f.InjectConventions(srv)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
func main() {
	cfg, _, ds, reset, install := setup()
	defer ds.DB.Close()

	if install && confirm("This will erase any data in the database. Confirm ?") {
		err := ds.Install()
		if err != nil {
			log.Fatalln("Unable to create the tables: " + err.Error())
		}
		log.Println("Tables created")
		os.Exit(0)
	} else if reset {
		err := ds.ResetRootAccount()
		if err != nil {
			log.Fatalln("Unable to reset the root account: " + err.Error())
		}
		log.Println("Root account reset. Don't forgot to delete it once logged")
		os.Exit(0)
	}

	puller := feeder.NewHTTPFeeder(cfg.Puller.URL, cfg.Puller.Login, cfg.Puller.Password, cfg.Puller.Promotions)
	period, err := time.ParseDuration(cfg.Puller.Period)
	if err != nil {
		log.Fatalln("Unable to start the convention feeder: " + err.Error())
	}
	startFeederDaemon(puller, ds, period)

	www := handler.NewService(ds)
	log.Println("Listening on " + cfg.HTTP.Host)
	www.Listen(cfg.HTTP.Host, cfg.HTTP.Certificate, cfg.HTTP.PrivateKey)

}
