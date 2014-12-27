package handler

/*package user

import (
	"log"
	"net/http"

	"github.com/fhermeni/wints/handler"
	"github.com/gorilla/mux"
)

func reportError(w http.ResponseWriter, msg string, e error) bool {
	if e != nil {
		switch e {
		case ErrNotFound:
			http.Error(w, e.Error(), http.StatusNotFound)
			return true
		case ErrExists:
			http.Error(w, e.Error(), http.StatusConflict)
			return true
		default:
			log.Printf("%s: %s\n", msg, e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return true
		}
	}
	return false
}

func Routes(us UserService) []handler.Route {
	return []handler.Route{
		handler.Route{"GET", "/profile", withService(Profile, us)},
		handler.Route{"PUT", "/profile", withService(SetProfile, us)},
		handler.Route{"PUT", "/profile/password", withService(SetPassword, us)},
		handler.Route{"POST", "/profile/password", withService(NewPassword, us)},
		handler.Route{"DELETE", "/profile/password", withService(ResetPassword, us)},
		handler.Route{"GET", "/users/", withService(Admins, us)},
		handler.Route{"POST", "/users/", withService(NewAdmin, us)},
		handler.Route{"POST", "/users/{email}/{role}", withService(SetRole, us)},
		handler.Route{"POST", "/login", withService(SetRole, us)},
		handler.Route{"POST", "/users/{email}", withService(Login, us)},
	}
}

func withService(cb func(UserService, http.ResponseWriter, *http.Request), us UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cb(us, w, r)
	}
}

func Profile(us UserService, w http.ResponseWriter, r *http.Request) {
	email, err := handler.Session(w, r)
	if err != nil {
		return
	}
	u, err := us.Get(email)
	if reportError(w, "Unable to get user '"+email+"'", err) {
		return
	}
	handler.JsonReply(w, false, u)
}

func SetProfile(us UserService, w http.ResponseWriter, r *http.Request) {
	email, err := handler.Session(w, r)
	if err != nil {
		return
	}

	var p User
	if handler.JsonRequest(w, r, &p) != nil {
		return
	}
	err = us.SetProfile(email, p.Firstname, p.Lastname, p.Tel)
	if reportError(w, "Unable to update the profile", err) {
		return
	}
	u, err := us.Get(email)
	if reportError(w, "Unable to get user '"+email+"'", err) {
		return
	}
	handler.JsonReply(w, false, u)
}

type PasswordRenewal struct {
	Old string
	New string
}

func SetPassword(us UserService, w http.ResponseWriter, r *http.Request) {
	email, err := handler.Session(w, r)
	if err != nil {
		return
	}

	var p PasswordRenewal
	if handler.JsonRequest(w, r, &p) != nil {
		return
	}
	err = us.SetPassword(email, []byte(p.Old), []byte(p.New))
	reportError(w, "Unable to change the password", err)
}

func ResetPassword(us UserService, w http.ResponseWriter, r *http.Request) {
	_, err := handler.Session(w, r)
	if err != nil {
		return
	}
	var email string
	if handler.JsonRequest(w, r, &email) != nil {
		return
	}
	_, err = us.ResetPassword(string(email))
	if reportError(w, "Unable to generate a token for a password reset", err) {
		return
	}
	//TODO: mail for token
	/*err = mailer.Mail([]string{string(email)}, "mails/account_reset.txt", struct {
		WWW   string
		Token string
	}{cfg.Http.Host, token})
	//reportError(w, "Unable to send the token for a password reset by mail", err)
}

func NewPassword(us UserService, w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	password := r.FormValue("password")
	email, err := us.NewPassword(token, []byte(password))
	if reportError(w, "Unable to create a new password from the token", err) {
		return
	}
	_, err = us.Register(email, password)
	if err != nil {
		http.Redirect(w, r, "/#badLogin", 302)
		return
	}
	handler.SetSession(email, w)
	http.Redirect(w, r, "/", 302)
}

func SetRole(us UserService, w http.ResponseWriter, r *http.Request) {
	_, err := handler.Session(w, r)
	if err != nil {
		return
	}
	target, _ := mux.Vars(r)["email"]
	var role string
	if handler.JsonRequest(w, r, &role) != nil {
		return
	}
	err = us.SetRole(target, string(role))
	if reportError(w, "", err) {
		return
	}
	//TODO: mail
	//mailer.Mail([]string{target}, "mails/priv_"+string(role)+".txt", nil)
}

func NewAdmin(us UserService, w http.ResponseWriter, r *http.Request) {
	_, err := handler.Session(w, r)
	if err != nil {
		return
	}
	var nu User
	if handler.JsonRequest(w, r, &nu) != nil {
		return
	}
	err = us.New(nu)
	if reportError(w, "Unable to create the user "+nu.Fullname(), err) {
		return
	}
}

func RmUser(us UserService, w http.ResponseWriter, r *http.Request) {
	target, _ := mux.Vars(r)["email"]
	err := us.Rm(target)
	reportError(w, "Unable to remove user '"+target+"'", err)
}

func Admins(us UserService, w http.ResponseWriter, r *http.Request) {
	users, err := us.List("tutor", "major", "admin", "roles")
	if err != nil {
		reportError(w, "Unable to list users", err)
	}
	handler.JsonReply(w, false, users)
}

func Login(us UserService, w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")
	_, err := us.Register(login, password)
	if err != nil {
		http.Redirect(w, r, "/#badLogin", 302)
		return
	}
	handler.SetSession(login, w)
	http.Redirect(w, r, "/", 302)
}*/
