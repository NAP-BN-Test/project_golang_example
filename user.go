package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "All User")
	fmt.Fprintf(w, "All User")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, name)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete user")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Update user")
}
