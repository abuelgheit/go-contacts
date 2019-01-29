package controllers

import (
	"encoding/json"
	"fmt"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateTodo = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	todo := &models.Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	todo.UserID = user
	resp := todo.CreateTodo()
	u.Respond(w, resp)
}

var GetTodos = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTodos(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTodo = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint)
	params := mux.Vars(r)
	todoID64, err := strconv.ParseUint(params["id"], 10, 32)
	todoID := uint(todoID64)
	if err != nil {
		u.Respond(w, u.Message(false, "Error"))
		return
	}
	data := models.GetTodo(todoID, userID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteTodo = func(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user").(uint)
	params := mux.Vars(r)
	todoID64, err := strconv.ParseUint(params["id"], 10, 32)
	todoID := uint(todoID64)
	if err != nil {
		u.Respond(w, u.Message(false, "Error"))
		return
	}
	data := models.DeleteTodo(todoID, userID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
	return
}
