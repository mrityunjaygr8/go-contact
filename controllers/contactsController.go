package controllers

import (
	"net/http"
	"go-contact/models"
	"encoding/json"
	u "go-contact/utils"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	contact_id := mux.Vars(r)["id"]
	contact_id_uid, err := strconv.Atoi(contact_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	contact.UserId = user
	data := models.UpdateContact(uint(contact_id_uid), user, contact)
	u.Respond(w, data)
}

var GetContact = func(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user").(uint)
	contact_id := mux.Vars(r)["id"]
	contact_id_uid, err := strconv.Atoi(contact_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := models.GetContact(uint(contact_id_uid), user_id)
	resp := u.Message(true, "success")
	if data == nil {
		extra := "Either this contact does not exist or you do not have access to this contact"
		resp["extra"] = extra
	}
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user").(uint)
	contact_id := mux.Vars(r)["id"]
	contact_id_uid, err := strconv.Atoi(contact_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp := models.DeleteContact(uint(contact_id_uid), user_id)
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}