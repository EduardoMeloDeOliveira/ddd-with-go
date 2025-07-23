package service

import (
	"errors"

	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/application/dto"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/entity"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

func (uService *UserService) SaveUser(registerRequest dto.UserRequestDTO) (*entity.User, error) {

	user := entity.NewUser(registerRequest.Name, registerRequest.Email)

	if err := uService.userRepository.Save(user); err != nil {
		return nil, err
	}

	return user, nil

}

func (uService *UserService) GetById(id string) (*entity.User, error) {
	user, err := uService.userRepository.FindById(id)

	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (uService *UserService) GetAll() ([]*entity.User, error) {
	users, err := uService.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uService *UserService) UpdateUser(userRequestDto dto.UserRequestDTO, id string) (*entity.User, error) {

	user, err := uService.GetById(id)

	if err != nil {
		return nil, err
	}

	newUser := entity.NewUser(userRequestDto.Name, userRequestDto.Email)
	newUser.ID = user.ID

	if err := uService.userRepository.Save(newUser); err != nil {
		return nil, err
	}

	return newUser, err
}

func (uService *UserService) Delete(id string) error {
	_, err := uService.GetById(id)

	if err != nil {
		return err
	}

	if err := uService.userRepository.Delete(id); err != nil {
		return err
	}

	return nil

}
