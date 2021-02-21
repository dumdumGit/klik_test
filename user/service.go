package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailExist(input EmailInput) (bool, error)
	GetUserById(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	theUser, err := s.repository.FindByEmail(email)
	if err != nil {
		return theUser, err
	}

	if theUser.Id == 0 {
		return theUser, errors.New("No Email not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(theUser.PasswordHash), []byte(password))
	if err != nil {
		return theUser, err
	}

	return theUser, nil
}

func (s *service) IsEmailExist(input EmailInput) (bool, error) {
	email := input.Email

	check, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if check.Id == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("No User Found")
	}

	return user, nil
}
