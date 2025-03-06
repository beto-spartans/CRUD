package http

import (
	"net/http"
)

func NewRouter(controller *UserController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetUser(w, r)
		case http.MethodPut:
			controller.UpdateUser(w, r)
		case http.MethodDelete:
			controller.DeleteUser(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.CreateUser(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
