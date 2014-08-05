package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/fhermeni/wints/backend"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"os"
	"errors"
	"strconv"
	"math/rand"
	"io/ioutil"
	"time"
)

const (
	ROOT_API = "/api/v1"
)

var DB *sql.DB

func reportIfError(w http.ResponseWriter, header string, err error) bool {
	if err == nil {
		return false
	}
	if be, ok := err.(*backend.BackendError); ok {
		http.Error(w, be.Error(), be.Status)
	} else {
		http.Error(w, "", http.StatusInternalServerError)
		log.Printf("%s: %s\n", header, err.Error())
	}
	return true
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := backend.ExtractEmail(r)
	if err != nil {
		http.ServeFile(w, r, "static/login.html")
	} else {
		http.Redirect(w, r, "/home", 302)
	}
}

func getLongDefenses(w http.ResponseWriter, r *http.Request) {
	cnt, err := backend.LongDefense(DB, "si/ifi")
	if !reportIfError(w, "", err) {
		w.Write([]byte(cnt))
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	n, err := backend.ExtractEmail(r)
	if reportIfError(w, "", err) {
		return
	}
	if len(n) == 0 {
		log.Printf("No cookie");
		http.Redirect(w, r, "/", 302)
	}
	u, err := backend.GetUser(DB, n)
	if !reportIfError(w, "", err) {
		if u.Role != "" {
			http.ServeFile(w, r, "static/admin.html")
			return
		}
		log.Printf("No dashboard available")
	}
}

func jsonReply(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	err := enc.Encode(j)
	reportIfError(w, "Error while serialising response", err)
}

func fileReply(w http.ResponseWriter, mime, filename string, cnt []byte) {
	w.Header().Set("Content-type", mime)
	w.Header().Set("Content-disposition", "attachment; filename=" + filename)
	_, err := w.Write(cnt)
	reportIfError(w, "", err)
}

func jsonRequest(w http.ResponseWriter, r *http.Request, j interface{}) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&j)
	if err != nil {
		log.Printf("Bad JSON message format: %s", err)
		http.Error(w, "Bad JSON message format", http.StatusBadRequest)
		return err
	}
	return nil
}

type PendingConventionMsg struct {
	C       backend.Convention
	Known   []backend.User
	Pending int
}

func RandomPendingConvention(w http.ResponseWriter, r *http.Request, email string) {
	cc, err := backend.GetRawConventions(DB)
	nb := len(cc)
	if nb == 0 {
		jsonReply(w, PendingConventionMsg{backend.Convention{}, nil, 0})
		return
	}
	c := cc[rand.Intn(nb)];
	known, err := backend.Admins(DB)
	if !reportIfError(w, "Unable to get the possible tutors", err) {
		jsonReply(w, PendingConventionMsg{c, known, nb})
	}
}

func GetAllConventions(w http.ResponseWriter, r *http.Request, email string) {
	conventions, err := backend.GetConventions(DB)
	if !reportIfError(w, "Unable to get the conventions ", err) {
		jsonReply(w, conventions)
	}
}


func GetAdmins(w http.ResponseWriter, r *http.Request, email string) {
	admins, err := backend.Admins(DB)
	if !reportIfError(w, "Unable to get the possible admins", err) {
		jsonReply(w, admins)
	}
}

type NewUser struct {
	Firstname string
	Lastname  string
	Tel       string
	Email     string
	Priv      string
}

func NewAdmin(w http.ResponseWriter, r *http.Request, email string) {
	var nu NewUser
	err := jsonRequest(w, r, &nu)
	if err != nil {
		return
	}
	p := backend.User{nu.Firstname, nu.Lastname, nu.Email, nu.Tel, nu.Priv}
	err = backend.NewUser(DB, p)
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, " User '"+nu.Email+"'("+nu.Priv+") created");
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	if target != email {
		http.Error(w, "You cannot get the profile of another person", http.StatusForbidden)
		return
	}
	u, err := backend.GetUser(DB, email)
	if !reportIfError(w, "", err) {
		jsonReply(w, u)
	}

}
func RmUser(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	err := backend.RmUser(DB, target)
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, " User '"+target+"' deleted");
	}
}

func GetDefense(w http.ResponseWriter, r *http.Request) {
	fmt := r.FormValue("fmt")
	if (len(fmt) == 0) {
		fmt = "embedded"
	}
	var cnt string
	var err error
	switch (fmt) {
	case "public":
		cnt, err = backend.LongDefense(DB, "si/ifi")
	case "embedded":
		email, _ := backend.ExtractEmail(r)
		if (len(email) == 0) {
			http.Redirect(w, r, "/", 302)
	    	return
		}
		u, _ := backend.GetUser(DB, email)
		if u.CompatibleRole("admin") {
			cnt, err = backend.Defense(DB, "si/ifi")
		} else {
			http.Error(w, "Unsufficient permission", http.StatusUnauthorized)
			return
		}
	}
	w.Write([]byte(cnt))
	reportIfError(w, "", err)
}

type DefensesPost struct {
	Short string
	Long  string
}

func PostDefense(w http.ResponseWriter, r *http.Request, email string) {
	var dp DefensesPost
	err := jsonRequest(w, r, &dp)
	if err != nil {
		return
	}
	log.Printf("%s\n", dp.Short)
	log.Printf("%s\n", dp.Long)

	err = backend.SaveDefense(DB, "si/ifi", string(dp.Short), string(dp.Long))
	reportIfError(w, "", err)

	/*cnt, err := ioutil.ReadAll(r.Body)

	//Separator short from long
	var f interface{}
	err = json.Unmarshal(cnt, &f)
	if reportIfError(w, "", err) {
		return
	}
    m := f.(map[string]interface{})

	if reportIfError(w, "", err) {
		return
	}
	log.Printf("%s\n",m["Short"]);
	log.Printf("%s\n",m["Long"]);
	//err = backend.SaveDefense(DB, "si/ifi", string(m["Short"].([]byte)), string(m["Long"].([]byte)))
	//reportIfError(w, "", err)*/
}

func Login(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	_, err := backend.Register(DB, login, password)
	if err != nil {
		http.Redirect(w, r, "/#badLogin", 302);
		return
	}
	backend.SetSession(login, w)
	backend.LogActionInfo(login, " log in");
	http.Redirect(w, r, "/", 302);
}

func Logout(w http.ResponseWriter, r *http.Request, email string) {
	backend.ClearSession(w);
	http.Redirect(w, r, "/", 302);
}

func CommitPendingConvention(w http.ResponseWriter, r *http.Request, email string) {
	var c backend.Convention
	err := jsonRequest(w, r, &c)
	if err != nil {
		return
	}
	tEmail := c.Tutor.Email
	_, err = backend.GetUser(DB, tEmail)
	if err != nil {
		err = backend.NewUser(DB, c.Tutor)
		if reportIfError(w, "Unable to register the tutor "+c.Tutor.String(), err) {
			return
		}
	}
	err = backend.RegisterInternship(DB, c, true);
	if reportIfError(w, "Unable to register the internship", err) {
		return
	}
	backend.LogActionInfo(email, "validate the convention of '"+c.Stu.P.Email+"'");
	backend.RescanPending(DB, c)
}

type PersonUpdate struct {
	Firstname string
	Lastname  string
	Tel       string
}

func ChangeProfile(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	if target != email {
		http.Error(w, "You cannot update the profile of another person", http.StatusForbidden)
		return
	}
	var p PersonUpdate
	if jsonRequest(w, r, &p) != nil {
		return
	}
	err := backend.SetProfile(DB, email, p.Firstname, p.Lastname, p.Tel)
	u, err := backend.GetUser(DB, email)
	if !reportIfError(w, "", err) {
		jsonReply(w, u)
	}
}

type PasswordRenewal struct {
	OldPassword string
	NewPassword string
}

func ChangePassword(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	if target != email {
		http.Error(w, "You cannot update the profile of another person", http.StatusForbidden)
		return
	}

	var p PasswordRenewal
	if jsonRequest(w, r, &p) != nil {
		return
	}

	err := backend.NewPassword(DB, email, []byte(p.OldPassword), []byte(p.NewPassword))
	reportIfError(w, "", err)
}

func UpdateMajor(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	m, err := requiredContent(w, r)
	if err != nil {
		return
	}
	err = backend.SetMajor(DB, target, string(m))
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, "Major of '"+target+"' set to '"+string(m)+"'");
	}
}

func MarkKind(cb func(http.ResponseWriter, *http.Request, string), kind string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cb(w, r, kind)
	}
}

func GetReport(kind string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target, _ := mux.Vars(r)["email"]
		cnt, err := backend.Report(DB, target, kind)
		if reportIfError(w, "", err) {
			return
		}

		u, _ := backend.GetUser(DB, target)

		fileReply(w, "application/pdf", u.Lastname + "-" + kind + ".pdf", cnt)
	}
}

func UpdateDeadline(kind string) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := requiredContent(w, r)
		if err != nil {
			return
		}
		t, err := time.Parse("2/1/2006", string(d))
		if reportIfError(w, "Invalid format", err) {
			return
		}
		target, _ := mux.Vars(r)["email"]

		//Other levels are ok
		err = backend.UpdateDeadline(DB, target, kind, t)
		reportIfError(w, "", err);
	}
}

func UploadReport(kind string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d, err := requiredContent(w, r)
		if err != nil {
			return
		}

		target, _ := mux.Vars(r)["email"]
		err = backend.SetReport(DB, target, kind, d)
		reportIfError(w, "", err);
	}
}

func UpdateMark(w http.ResponseWriter, r *http.Request, kind string) {
	d, err := requiredContent(w, r)
	if err != nil {
		return
	}
	i, err := strconv.Atoi(string(d))
	if reportIfError(w, "Invalid format", err) {
		return
	}
	if i < 0 || i > 20 {
		reportIfError(w, "Mark must be between 0 and 20", errors.New("Mark must be between 0 and 20"))
		return
	}
	target, _ := mux.Vars(r)["email"]
	//Other levels are ok
	err = backend.UpdateGrade(DB, target, kind, i)
	reportIfError(w, "", err);
}

func myStudent(w http.ResponseWriter, mine, target string) bool {
	u, err := backend.GetUser(DB, target);
	if err != nil {
		reportIfError(w, "", err)
		return false
	}
	if u.Role == "" {
		c, err := backend.GetConvention(DB, target)
		if reportIfError(w, "", err) {
			return false
		}
		if c.Tutor.Email != mine {
			http.Error(w, "The student is not under your responsibility", http.StatusForbidden)
			return false
		}
	}
	return true
}

func requiredContent(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	cnt, err := ioutil.ReadAll(r.Body)
	if reportIfError(w, "", err) {
		return []byte{}, err
	}
	if len(cnt) == 0 {
		reportIfError(w, "", errors.New("Missing request body"))
	}
	return cnt, nil
}

func GrantRole(w http.ResponseWriter, r *http.Request, email string) {
	target, _ := mux.Vars(r)["email"]
	role, err := requiredContent(w, r)
	if err != nil {
		return
	}
	err = backend.GrantPrivilege(DB, target, string(role))
	if !reportIfError(w, "", err) {
		backend.LogActionInfo(email, "'"+target+"' granted to role '"+string(role)+"'");
	}
}

func RequireToken(cb func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n, _ := backend.ExtractEmail(r)
		if (len(n) == 0) {
			http.Redirect(w, r, "/", 302)
			return
		}
		cb(w, r, n)
	}
}

func RequireRole(cb func(http.ResponseWriter, *http.Request, string), role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, _ := backend.ExtractEmail(r)
		if (len(email) == 0) {
			http.Redirect(w, r, "/", 302)
			return
		}

		u, _ := backend.GetUser(DB, email)
		if u.CompatibleRole(role) {
			cb(w, r, email)
		} else {
			http.Error(w, "Unsufficient permission", http.StatusUnauthorized)
		}
	}
}

func migrateReports() {
	rows, err := DB.Query("select student,midtermDeadline from internships")
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	defer rows.Close();
	for rows.Next() {
		var m string
		var d time.Time
		rows.Scan(&m, &d)
		_, err = backend.NewReport(DB, m, "midterm", d)
		if err != nil {
			log.Printf("midterm: %s\n", err)
		}
		t, _ := time.Parse("02/01/2006", "01/09/2014")
		_, err = backend.NewReport(DB, m, "final", t)
		if err != nil {
			log.Printf("final: %s\n", err)
		}

		_, err = backend.NewReport(DB, m, "supReport", t)
		if err != nil {
			log.Printf("final: %s\n", err)
		}

	}
}
func main() {

	cfg, err := backend.ReadConfig("./wints.conf")
	if err != nil {
		log.Fatalln(err.Error())
	}

	backend.InitLogger(cfg.Logfile)
	defer backend.CloseLogger();
	dbUrl := os.Getenv("DATABASE_URL")

	if len(dbUrl) == 0 {
		dbUrl = cfg.DBurl
	}
	DB, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, v := range os.Args {
		if v == "-i" {
			err := backend.NewUser(DB, backend.User{"R", "oot", "root@localhost", "", "root"})
			if err != nil {
				backend.LogError("Unable to create the root account: " + err.Error());
			} else {
				backend.LogInfo("Root account created");
			}
			return
		} else if v == "-s" {
			delay, err := time.ParseDuration(cfg.RefreshPeriod)
			if err != nil {
				log.Fatalln(err.Error())
			}
			backend.DaemonConventionsPuller(DB, cfg.WWWConventionsURL,
				cfg.WWWConventionsLogin,
				cfg.WWWConventionsPassword,
				delay);
		}
	}

	//Rest stuff
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/home", homeHandler)

	//migrateReports()
	r.HandleFunc(ROOT_API+"/pending/_random", RequireRole(RandomPendingConvention, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API+"/conventions/", RequireRole(GetAllConventions, "major")).Methods("GET")
	r.HandleFunc(ROOT_API+"/conventions/", RequireRole(CommitPendingConvention, "admin")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/major", RequireRole(UpdateMajor, "major")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/midterm/deadline", UpdateDeadline("midterm")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/midterm/report", UploadReport("midterm")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/midterm/report", GetReport("midterm")).Methods("GET")
	r.HandleFunc(ROOT_API+"/conventions/{email}/midterm/mark", MarkKind(UpdateMark, "midterm")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/final/deadline", UpdateDeadline("final")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/final/report", UploadReport("final")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/final/report", GetReport("final")).Methods("GET")
	r.HandleFunc(ROOT_API+"/conventions/{email}/final/mark", MarkKind(UpdateMark, "final")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/supervisor/deadline", UpdateDeadline("supReport")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/supervisor/report", UploadReport("major")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/supervisor/mark", MarkKind(UpdateMark, "supReport")).Methods("POST")
	r.HandleFunc(ROOT_API+"/conventions/{email}/supervisor/report", GetReport("supReport")).Methods("GET")


	r.HandleFunc(ROOT_API+"/users/", RequireRole(GetAdmins, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API+"/users/", RequireRole(NewAdmin, "root")).Methods("POST")
	r.HandleFunc(ROOT_API+"/users/{email}/role", RequireRole(GrantRole, "root")).Methods("POST")
	r.HandleFunc(ROOT_API+"/users/{email}", RequireRole(RmUser, "root")).Methods("DELETE")
	r.HandleFunc(ROOT_API+"/users/{email}", RequireToken(GetUser)).Methods("GET")
	r.HandleFunc(ROOT_API+"/users/{email}/password", RequireToken(ChangePassword)).Methods("POST")
	r.HandleFunc(ROOT_API+"/users/{email}/", RequireToken(ChangeProfile)).Methods("POST")
	r.HandleFunc(ROOT_API+"/login", Login).Methods("POST")
	r.HandleFunc(ROOT_API+"/logout", RequireToken(Logout)).Methods("GET")
	//r.HandleFunc(ROOT_API+"/defenses", RequireRole(GetDefense, "admin")).Methods("GET")
	r.HandleFunc(ROOT_API+"/defenses", GetDefense).Methods("GET")
	r.HandleFunc(ROOT_API+"/defenses", RequireRole(PostDefense, "admin")).Methods("POST")
	//r.HandleFunc("/defenses/program", getDefensesProgram).Methods("GET")
	fs := http.Dir("static/")
	fileHandler := http.FileServer(fs)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", r)
	log.Println("Daemon started")
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	defer DB.Close()
}
