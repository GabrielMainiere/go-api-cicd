package main

import (
	"api-go/user"
	"log"
	"net/http"
	"os"
)

func main() {
	repo, err := user.NewUserRepositoryFromEnv()
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	service := user.NewUserService(repo)
	controller := user.NewUserController(service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

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

	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
