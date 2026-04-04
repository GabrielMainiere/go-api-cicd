package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	service *UserService
}

func NewUserController(s *UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path)

	user, err := c.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	created, err := c.service.CreateUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(created); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path)

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updated, err := c.service.UpdateUser(id, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path)

	err := c.service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func extractID(path string) int {
	parts := strings.Split(path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	return id
}