package controllers

import (
	"crud/models"
	"fmt"
	"net/http"
	"text/template"
)

func ReadUsers(w http.ResponseWriter, r *http.Request) {

	view, err := template.ParseFiles("views/readUsers.html", "views/layout.html")
	if err != nil {
		fmt.Print("Read user view error")
	}
	usuarios := models.ReadUsers()

	view.Execute(w, usuarios)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		email := r.FormValue("email")
		password := r.FormValue("password")

		models.CreateUser(nombre, email, password)
		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	view, err := template.ParseFiles("views/createUser.html", "views/layout.html")
	if err != nil {
		fmt.Print("create user view error")
	}
	view.Execute(w, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteUser(id)
	http.Redirect(w, r, "/users/", http.StatusFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		email := r.FormValue("email")
		password := r.FormValue("password")

		models.UpdateUser(id, nombre, email, password)
		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	view, err := template.ParseFiles("views/UpdateUser.html", "views/layout.html")
	if err != nil {
		fmt.Print("create user view error")
	}
	usuario := models.ReadUser(id)

	view.Execute(w, usuario)

}
