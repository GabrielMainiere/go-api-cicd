package user

import (
	"errors"
)

type InMemoryUserRepository struct {
	users  []User
	nextID int
}

func NewUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users:  []User{},
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepository) GetByID(id int) (*User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) Create(u User) (*User, error) {
	u.ID = r.nextID
	r.nextID++
	r.users = append(r.users, u)
	return &u, nil
}

func (r *InMemoryUserRepository) Update(id int, u User) (*User, error) {
	for i, user := range r.users {
		if user.ID == id {
			u.ID = id
			r.users[i] = u
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) Delete(id int) error {
	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}