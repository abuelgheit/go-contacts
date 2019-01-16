package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/co/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/co/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/co/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/co/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
	router.HandleFunc("/co/api/me/contacts/{id}", controllers.GetContact).Methods("GET")
	router.HandleFunc("/co/api/me/contacts/{id}", controllers.DeleteContact).Methods("DELETE")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
