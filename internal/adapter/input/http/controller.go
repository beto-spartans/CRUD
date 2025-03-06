package http

import (
	"CRUD/internal/usecase/user"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type UserController struct {
	service user.UserServiceInterface
}

func NewUserController(service user.UserServiceInterface) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()
	newUser, err := c.service.CreateUser(id, request.Name, request.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userIDStr := strings.TrimPrefix(path, "/users/")
	if userIDStr == "" {
		http.Error(w, "ID no proporcionado", http.StatusBadRequest)
		return
	}
	user, err := c.service.GetUser(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userIDStr := strings.TrimPrefix(path, "/users/")
	if userIDStr == "" {
		http.Error(w, "ID no proporcionado", http.StatusBadRequest)
		return
	}

	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	updatedUser, err := c.service.UpdateUser(userIDStr, request.Name, request.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userIDStr := strings.TrimPrefix(path, "/users/")
	if userIDStr == "" {
		http.Error(w, "ID no proporcionado", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteUser(userIDStr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
