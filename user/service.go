package user

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(u User) (*User, error) {
	return s.repo.Create(u)
}

func (s *UserService) UpdateUser(id int, u User) (*User, error) {
	return s.repo.Update(id, u)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}