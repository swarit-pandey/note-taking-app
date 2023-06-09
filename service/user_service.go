package service

import (
	"errors"

	"github.com/sprectza/note-taking-app.git/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user repository.User) error
	LoginUser(email, password string) (repository.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) RegisterUser(user repository.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepository.CreateUser(user)
}

func (s *userService) LoginUser(email, password string) (repository.User, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}
