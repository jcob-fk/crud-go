package routes

import (
	"crud/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", controllers.ReadUsers).Methods("GET")
	userRouter.HandleFunc("/create", controllers.CreateUser).Methods("GET", "POST")
	userRouter.HandleFunc("/delete", controllers.DeleteUser).Methods("GET")
	userRouter.HandleFunc("/update", controllers.UpdateUser).Methods("GET", "POST")
	return router
}
