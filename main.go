package main

import (
	"net/http"
	"log"
	"api-go/user"
)

func main() {
	repo := user.NewUserRepository()
	service := user.NewUserService(repo)
	controller := user.NewUserController(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.ListUsers(w, r)
		case http.MethodPost:
			controller.CreateUser(w, r)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetUser(w, r)
		case http.MethodPut:
			controller.UpdateUser(w, r)
		case http.MethodDelete:
			controller.DeleteUser(w, r)
		}
	})

	log.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3001", mux))
}