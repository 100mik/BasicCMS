package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success bool `json:"success"`
	Message string	`json:"message"`
	Data interface{}	`json:"data"`
}

func ResponseWriter(flag bool, msg string, data interface{}, status int, w http.ResponseWriter) {
	response := Response{
		Success: flag,
		Message: msg,
		Data:    data,
	}

	payload, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}

func GetCurrUser(r *http.Request) (CurrUser, error) {
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error in reading cookie\n\t" + err.Error())
		return CurrUser{}, err
	}
	value := CurrUser{}
	if err = CookieHandler.Decode("session", c.Value, &value); err == nil {
		return value, nil
	} else {
		fmt.Println("Could not read cookie data\n" + err.Error())
		return CurrUser{}, err
	}
}
