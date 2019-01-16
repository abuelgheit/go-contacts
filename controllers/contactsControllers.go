package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
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

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetContact = func(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user").(uint)
	params := mux.Vars(r)
	contactID64, err := strconv.ParseUint(params["id"], 10, 32)
	contactID := uint(contactID64)
	if err != nil {
		u.Respond(w, u.Message(false, "Error"))
		return
	}
	data := models.GetContact(contactID, userID)
	if data != nil {
		resp := u.Message(true, "success")
		resp["data"] = data
		u.Respond(w, resp)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	resp := u.Message(false, "Not Found")
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint)
	params := mux.Vars(r)
	contactID64, err := strconv.ParseUint(params["id"], 10, 32)
	contactID := uint(contactID64)
	if err != nil {
		u.Respond(w, u.Message(false, "Error"))
		return
	}
	data, deletedContact := models.DeleteContact(contactID, userID)
	if deletedContact != nil {
		resMesage := "contact " + deletedContact.Name + " deleted"
		resp := u.Message(true, resMesage)
		resp["data"] = data
		u.Respond(w, resp)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	resp := u.Message(false, "Not Found")
	u.Respond(w, resp)
}
