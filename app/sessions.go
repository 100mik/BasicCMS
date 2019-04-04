package app

import (
	"net/http"
)

func SetSession(w http.ResponseWriter, currUser CurrUser, age int) error {

	if encoded, err := CookieHandler.Encode("session", currUser); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: age,
		}
		http.SetCookie(w, cookie)
	} else {
		return err
	}
	return nil
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
