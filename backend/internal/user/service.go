package user

type Service interface {
    CreateUser(user *User) error
    GetUserByID(id uint) (*User, error)
    GetUserByEmail(email string) (*User, error)
    GetAllUsers() ([]User, error)
    UpdateUser(user *User) error
    DeleteUser(id uint) error
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo}
}

func (s *service) CreateUser(user *User) error {
    return s.repo.CreateUser(user)
}

func (s *service) GetUserByID(id uint) (*User, error) {
    return s.repo.GetUserByID(id)
}

func (s *service) GetUserByEmail(email string) (*User, error) {
    return s.repo.GetUserByEmail(email)
}

func (s *service) GetAllUsers() ([]User, error) {
    return s.repo.GetAllUsers()
}

func (s *service) UpdateUser(user *User) error {
    return s.repo.UpdateUser(user)
}

func (s *service) DeleteUser(id uint) error {
    return s.repo.DeleteUser(id)
}
