package user

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserRepository interface {
    GetAll() ([]User, error)
    GetByID(id int) (*User, error)
    Create(u User) (*User, error)
    Update(id int, u User) (*User, error)
    Delete(id int) error
}