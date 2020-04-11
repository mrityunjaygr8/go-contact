package controllers

import (
	"net/http"
	u "go-contact/utils"
	"go-contact/models"
	"encoding/json"
	"fmt"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		fmt.Println(string(err.Error()))
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}