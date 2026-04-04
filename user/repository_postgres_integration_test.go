//go:build integration

package user

import (
	"os"
	"testing"
)

func TestPostgresUserRepository_CRUD(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("RUN_INTEGRATION_TESTS is not true")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("DATABASE_URL not set")
	}

	repoInterface, err := NewPostgresUserRepository(databaseURL)
	if err != nil {
		t.Fatalf("failed to create postgres repository: %v", err)
	}

	repo, ok := repoInterface.(*PostgresUserRepository)
	if !ok {
		t.Fatal("repository is not PostgresUserRepository")
	}

	_, err = repo.db.Exec(`DELETE FROM users`)
	if err != nil {
		t.Fatalf("failed to cleanup users table: %v", err)
	}

	created, err := repo.Create(User{Name: "Bob", Email: "bob@mail.com"})
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	fetched, err := repo.GetByID(created.ID)
	if err != nil {
		t.Fatalf("get by id failed: %v", err)
	}
	if fetched.Email != "bob@mail.com" {
		t.Fatalf("expected email bob@mail.com, got %s", fetched.Email)
	}

	updated, err := repo.Update(created.ID, User{Name: "Bob Updated", Email: "bob.updated@mail.com"})
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}
	if updated.Name != "Bob Updated" {
		t.Fatalf("expected updated name Bob Updated, got %s", updated.Name)
	}

	allUsers, err := repo.GetAll()
	if err != nil {
		t.Fatalf("get all failed: %v", err)
	}
	if len(allUsers) != 1 {
		t.Fatalf("expected 1 user, got %d", len(allUsers))
	}

	if err := repo.Delete(created.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
}
