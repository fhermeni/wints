package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/fhermeni/wints/Godeps/_workspace/src/github.com/BurntSushi/toml"
	_ "github.com/fhermeni/wints/Godeps/_workspace/src/github.com/lib/pq"
	"github.com/fhermeni/wints/config"
	"github.com/fhermeni/wints/feeder"
	"github.com/fhermeni/wints/httpd"
	"github.com/fhermeni/wints/journal"
	"github.com/fhermeni/wints/mail"
	"github.com/fhermeni/wints/notifier"
	"github.com/fhermeni/wints/schema"
	"github.com/fhermeni/wints/sqlstore"
)

var cfg config.Config
var not *notifier.Notifier
var store *sqlstore.Store
var mailer mail.Mailer

func confirm(msg string) bool {
	os.Stdout.WriteString(msg + " (y/n) ?")
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	ret := string(b) == "y"
	return ret
}

func inviteRoot(em string) {
	p := schema.Person{Firstname: "root", Lastname: "root", Email: em, Tel: "n/a"}
	token, err := store.NewUser(p, schema.ROOT)
	not.Fatalln("Create root account", err)
	if err := not.InviteRoot(schema.User{Person: p}, p, string(token), err); err != nil {
		//Here, we delete the account as the root account was not aware of the creation
		store.RmUser(p.Email)
		not.Fatalln("Invite root", err)
	}
}

func newConventionReader(cfg config.Feeder, not *notifier.Notifier) feeder.ConventionReader {
	reader := feeder.NewHTTPConventionReader(cfg.URL, cfg.Login, cfg.Password, not.Log)
	reader.Encoding = cfg.Encoding
	return reader
}

func newMailer(fake bool) mail.Mailer {
	if fake {
		return &mail.Fake{WWW: cfg.HTTPd.WWW, Config: cfg.Mailer}
	}
	m, err := mail.NewSMTP(cfg.Mailer, cfg.HTTPd.WWW)
	if err != nil {
		log.Fatalln("SMTP Mailer: " + err.Error())
	}
	return m
}

func newJournal() journal.Journal {
	j, err := journal.FileBacked(cfg.Journal.Path)
	if err != nil {
		log.Fatalln("Unable to create logs: " + err.Error())
	}
	return j
}

func newStore() *sqlstore.Store {
	DB, err := sql.Open("postgres", cfg.Db.ConnectionString)
	not.Fatalln("Database connexion", err)
	store, _ := sqlstore.NewStore(DB, cfg.Internships)
	return store
}

func newFeeder(not *notifier.Notifier) feeder.Conventions {
	r := feeder.NewHTTPConventionReader(cfg.Feeder.URL, cfg.Feeder.Login, cfg.Feeder.Password, not.Log)
	r.Encoding = cfg.Feeder.Encoding
	f := feeder.NewCsvConventions(r, cfg.Feeder.Promotions, not.Log)
	return f
}

/*func runSpies() {

	if len(cfg.Spies) == 0 {
		log.Println("No spies to schedule")
	}
	cr := cron.New()
	for _, s := range cfg.Spies {
		switch strings.ToLower(s.Kind) {
		case "tutor":
			reviews := make(map[string]time.Duration)
			for _, r := range cfg.Internships.Reports {
				reviews[r.Kind] = r.Review.Duration
			}
			cr.AddFunc(s.Cron, spy.ReminderReporter)
		default:
			log.Fatalln("Unsupported spy '" + s.Kind + "'")
		}
		log.Println("Spy '" + s.Kind + "' scheduled")
	}
	cr.Start()
}*/

func install() {
	if !confirm("This will erase any data. Confirm ") {
		os.Exit(1)
	}
	err := store.Install()
	not.Fatalln("Creating tables", err)
}
func main() {

	makeRoot := flag.String("new-root", "", "Invite a root user")
	fakeMailer := flag.Bool("fake-mailer", false, "Don't send emails. Print them out stdout")
	installStore := flag.Bool("install-db", false, "install the database")
	conf := flag.String("conf", "wints.conf", "Wints configuration file")
	flag.Parse()

	if _, err := toml.DecodeFile(*conf, &cfg); err != nil {
		log.Fatalln(err.Error())
	}

	mailer = newMailer(*fakeMailer)
	not = notifier.New(mailer, newJournal())
	not.Log.Wipe()
	store = newStore()
	if *installStore {
		install()
		os.Exit(0)
	}

	_, err := store.Internships()
	not.Fatalln("Database communication", err)

	if len(*makeRoot) > 0 {
		inviteRoot(*makeRoot)
		os.Exit(0)
	}

	//runSpies()
	conventions := newFeeder(not)
	not.Println("Listening on "+cfg.HTTPd.WWW, nil)
	httpd := httpd.NewHTTPd(not, store, conventions, cfg.HTTPd, cfg.Internships)
	not.Fatalln("%s\n", httpd.Listen())
}
