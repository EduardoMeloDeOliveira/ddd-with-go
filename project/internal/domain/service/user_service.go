package service

import (
	"errors"

	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/entity"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (uService *UserService) SaveUser(name string, email string) (*entity.User, error) {

	if name == "" || email == "" {
		return nil, errors.New("name and email can't be empty")
	}

	user := entity.NewUser(name, email)

	if err := uService.userRepo.Save(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (UserService *UserService) GetById(id string) (*entity.User, error) {
	user, err := UserService.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (UserService *UserService) GetAll() ([]*entity.User, error) {
	users, err := UserService.userRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
