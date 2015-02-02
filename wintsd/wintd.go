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

	_ "net/http/pprof"

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

func newServices(cfg config.Config) (mail.Mailer, *datastore.Service) {

	mailer, err := mail.NewSMTP(cfg.Mailer.Server, cfg.Mailer.Login, cfg.Mailer.Password, cfg.Mailer.Sender, cfg.HTTP.WWW, cfg.Mailer.Path)
	if err != nil {
		log.Fatalln("Unable to setup the mailer: " + err.Error())
	}

	DB, err := sql.Open("postgres", cfg.DB.URL)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}
	ds, err := datastore.NewService(DB, cfg.Reports, cfg.Majors)
	if err != nil {
		log.Fatalln("Unable connect to the database: " + err.Error())
	}
	return mailer, ds
}

func test(msg string, err error) bool {
	if err != nil {
		log.Println(msg + ": [FAIL]\n\t" + err.Error())
	} else {
		log.Println(msg + ": [OK]")
	}
	return err == nil
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

func newRoot(em string, ds *datastore.Service, m mail.Mailer) {
	u := internship.User{Firstname: "Root", Lastname: em, Email: em, Tel: "666"}
	tok, err := ds.NewTutor(u)
	if err != nil {
		log.Fatalf("Unable to create an account for '%s': %s\n", em, err.Error())
	}
	if err := ds.SetUserRole(em, internship.ROOT); err != nil {
		ds.RmUser(em)
		log.Fatalf("Unable to grant root privileges for '%s': %s\n", em, err.Error())
	}
	m.SendAdminInvitation(u, tok)
}

func main() {
	fakeMailer := flag.Bool("fakeMailer", false, "Use a fake mailer that print mail on stdout")
	cfgPath := flag.String("conf", "./wints.conf", "daemon configuration file")
	install := flag.Bool("install", false, "/!\\ Create database tables")
	makeRoot := flag.String("invite-root", "", "Invite a root user")
	testMail := flag.Bool("test-mailer", false, "Test the mailer by sending a mail to the administrator")
	testFeeder := flag.Bool("test-feeder", false, "Test the convention feeder")
	testAll := flag.Bool("test", false, "equivalent to --testMail --testDB --testFeeder")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalln("Error while parsing " + *cfgPath + ": " + err.Error())
	}
	log.Println("Parsing " + *cfgPath + ": OK")

	mailer, ds := newServices(cfg)
	defer ds.DB.Close()

	mailer.Fake(*fakeMailer)
	puller := feeder.NewHTTPFeeder(cfg.Puller.URL, cfg.Puller.Login, cfg.Puller.Password, cfg.Puller.Promotions, cfg.Puller.Encoding)

	if *install && confirm("This will erase any data in the database. Confirm ?") {
		err := ds.Install()
		if err != nil {
			log.Fatalln("Unable to create the tables: " + err.Error())
		}
		log.Println("Tables created")
		os.Exit(0)
	}

	if len(*makeRoot) > 0 {
		newRoot(*makeRoot, ds, mailer)
		os.Exit(0)
	}
	//Test connection anyway
	_, e := ds.Internships()
	ok := test("Database connection", e)

	if *testAll {
		*testFeeder = true
		*testMail = true
	}
	if *testMail {
		k := test("Mailing system", mailer.SendTest(cfg.Mailer.Sender))
		ok = ok && k
	}
	if *testFeeder {
		k := test("Convention feeder", puller.Test())
		ok = ok && k
	}
	if !ok {
		os.Exit(1)
	}

	period, err := time.ParseDuration(cfg.Puller.Period)
	if err != nil {
		log.Fatalln("Unable to start the convention feeder: " + err.Error())
	}
	startFeederDaemon(puller, ds, period)

	www := handler.NewService(ds, mailer, cfg.HTTP.Path)
	log.Println("Listening on " + cfg.HTTP.Listen)
	err = www.Listen(cfg.HTTP.Listen, cfg.HTTP.Certificate, cfg.HTTP.PrivateKey)
	if err != nil {
		log.Fatalln("Server exited:" + err.Error())
	}
}
