package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_CreateUser(t *testing.T) {
	repo := &fakeUserRepository{users: []User{}}
	service := NewUserService(repo)
	controller := NewUserController(service)

	payload := User{Name: "Carol", Email: "carol@mail.com"}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	controller.CreateUser(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}

	var created User
	if err := json.Unmarshal(rr.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if created.ID == 0 {
		t.Fatal("expected created user id to be greater than zero")
	}
	if created.Name != payload.Name {
		t.Fatalf("expected name %s, got %s", payload.Name, created.Name)
	}
}
