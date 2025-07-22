package repository

import "github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindById(id string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
}
