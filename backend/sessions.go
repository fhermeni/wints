package backend

import (
	"net/http"
)

func SetSession(email string, response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:  "session",
		Value: email,
		Path:  "/",
	}
	http.SetCookie(response, cookie)
}

func ExtractEmail(request *http.Request) (string, error) {
	var err error
	cookie, err := request.Cookie("session")
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
