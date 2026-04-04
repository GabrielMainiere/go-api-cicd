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
	users, _ := c.service.ListUsers()
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path)

	user, err := c.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	created, _ := c.service.CreateUser(u)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := extractID(r.URL.Path)

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	updated, err := c.service.UpdateUser(id, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updated)
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