package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(controller *UserController) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	return router
}
