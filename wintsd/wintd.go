//launch wints daemon
package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/davecheney/profile"

	"github.com/fhermeni/wints/cache"
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/datastore"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/handler"
	"github.com/fhermeni/wints/internship"
	"github.com/fhermeni/wints/journal"
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

func newMailer(cfg config.Config, j journal.Journal, fake bool) mail.Mailer {
	mailer, err := mail.NewSMTP(cfg.Mailer.Server, cfg.Mailer.Login, cfg.Mailer.Password, cfg.Mailer.Sender, cfg.HTTP.WWW, cfg.Mailer.Path, j)
	if err != nil {
		log.Fatalln("Unable to setup the mailer: " + err.Error())
	}
	mailer.Fake(fake)
	return mailer
}

func newJournal(cfg config.Config) journal.Journal {
	j, err := journal.FileBacked(cfg.Journal.Path)
	if err != nil {
		log.Fatalln("Unable to create logs: " + err.Error())
	}
	return j
}

func newServices(cfg config.Config) *datastore.Service {

	DB, err := sql.Open("postgres", cfg.DB.URL)
	if err != nil {
		log.Fatalln("Unable to connect to the Database: " + err.Error())
	}
	ds, err := datastore.NewService(DB, cfg.Reports, cfg.Surveys, cfg.Majors)
	if err != nil {
		log.Fatalln("Unable connect to the database: " + err.Error())
	}
	return ds
}

func newFeeder(cfg config.Config, j journal.Journal) feeder.Feeder {
	return feeder.NewHTTPFeeder(j, cfg.Puller.URL, cfg.Puller.Login, cfg.Puller.Password, cfg.Puller.Promotions, cfg.Puller.Encoding)
}

func newCache(backend internship.Service) *cache.Cache {
	cc, err := cache.NewCache(backend)
	if err != nil {
		log.Fatalln("Unable to initiate the cache: " + err.Error())
	}
	return cc
}

func test(j journal.Journal, msg string, err error) bool {
	j.Log("wintsd", msg+" ", err)
	if err != nil {
		log.Fatalln(msg + ": " + err.Error())
	}
	return err == nil
}

func startFeederDaemon(j journal.Journal, f feeder.Feeder, srv internship.Service, frequency time.Duration) {
	j.Log("wintsd", "Scanning conventions", nil)
	go f.InjectConventions(srv)
	ticker := time.NewTicker(time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				j.Log("wintsd", "Scanning conventions", nil)
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

	//Set the number of procs
	runtime.GOMAXPROCS(runtime.NumCPU())

	fakeMailer := flag.Bool("fake-mailer", false, "Use a fake mailer that print mail on stdout")
	cfgPath := flag.String("conf", "./wints.conf", "daemon configuration file")
	install := flag.Bool("install", false, "/!\\ Create database tables")
	makeRoot := flag.String("invite-root", "", "Invite a root user")
	testMail := flag.Bool("test-mailer", false, "Test the mailer by sending a mail to the administrator")
	testFeeder := flag.Bool("test-feeder", false, "Test the convention feeder")
	testAll := flag.Bool("test", false, "equivalent to --test-mailer --test-feeder")
	blankConf := flag.Bool("generate-config", false, "Print a default configuration file")
	cpuProfile := flag.Bool("cpu-profile", false, "Profile CPU usage")
	flag.Parse()

	if *blankConf {
		config.Blank()
		os.Exit(0)
	}

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalln("Error while parsing " + *cfgPath + ": " + err.Error())
	}

	if *cpuProfile {
		prof := &profile.Config{
			CPUProfile:  true,
			ProfilePath: cfg.Journal.Path,
		}
		defer profile.Start(prof).Stop()
	}

	j := newJournal(cfg)
	mailer := newMailer(cfg, j, *fakeMailer)
	puller := newFeeder(cfg, j)
	ds := newServices(cfg)
	defer ds.DB.Close()

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

	cc := newCache(ds)
	//Test connection anyway
	_, e := cc.Internships()
	ok := test(j, "Database connection", e)

	if *testAll {
		*testFeeder = true
		*testMail = true
	}
	if *testMail {
		k := test(j, "Mailing system", mailer.SendTest(cfg.Mailer.Sender))
		ok = ok && k
	}
	if *testFeeder {
		k := test(j, "Convention feeder", puller.Test())
		ok = ok && k
	}
	if !ok {
		os.Exit(1)
	}

	period, err := time.ParseDuration(cfg.Puller.Period)
	if err != nil {
		log.Fatalln("Unable to start the convention feeder: " + err.Error())
	}

	startFeederDaemon(j, puller, cc, period)

	j.Log("wintsd", "Starting wints", nil)
	www := handler.NewService(cc, mailer, cfg.HTTP.Path, j)
	j.Log("wintsd", "Listening on "+cfg.HTTP.Listen, nil)
	j.Log("wintsd", "Working over "+strconv.Itoa(runtime.NumCPU())+" CPU(s)", nil)
	log.Println("Wintsd running. See logs in folder '" + cfg.Journal.Path + "'")
	err = www.Listen(cfg.HTTP.Listen, cfg.HTTP.Certificate, cfg.HTTP.PrivateKey)
	if err != nil {
		log.Fatalln("Server exited:" + err.Error())
	}
}
