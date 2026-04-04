package user

import (
	"errors"
	"testing"
)

type fakeUserRepository struct {
	users []User
}

func (f *fakeUserRepository) GetAll() ([]User, error) {
	return f.users, nil
}

func (f *fakeUserRepository) GetByID(id int) (*User, error) {
	for _, u := range f.users {
		if u.ID == id {
			user := u
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (f *fakeUserRepository) Create(u User) (*User, error) {
	u.ID = len(f.users) + 1
	f.users = append(f.users, u)
	return &u, nil
}

func (f *fakeUserRepository) Update(id int, u User) (*User, error) {
	for i := range f.users {
		if f.users[i].ID == id {
			u.ID = id
			f.users[i] = u
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (f *fakeUserRepository) Delete(id int) error {
	for i := range f.users {
		if f.users[i].ID == id {
			f.users = append(f.users[:i], f.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func TestUserService_CreateAndList(t *testing.T) {
	repo := &fakeUserRepository{users: []User{}}
	service := NewUserService(repo)

	_, err := service.CreateUser(User{Name: "Alice", Email: "alice@mail.com"})
	if err != nil {
		t.Fatalf("create user failed: %v", err)
	}

	users, err := service.ListUsers()
	if err != nil {
		t.Fatalf("list users failed: %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
	if users[0].Name != "Alice" {
		t.Fatalf("expected user name Alice, got %s", users[0].Name)
	}
}
