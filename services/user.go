package services

import (
	"github.com/realjv3/gotasks/domain"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *domain.User) (*domain.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(password)

	return s.repo.Create(user)
}

func (s *userService) GetUser(id int) (*domain.User, error) {
	return s.repo.Get(id)
}
