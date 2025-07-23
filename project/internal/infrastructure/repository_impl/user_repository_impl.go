package repository_impl

import (
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repoImpl *UserRepositoryImpl) Save(user *entity.User) error {
	return repoImpl.db.Save(user).Error
}

func (repoImpl *UserRepositoryImpl) FindById(id string) (*entity.User, error) {
	var user entity.User

	err := repoImpl.db.First(&user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repoImpl *UserRepositoryImpl) FindAll() ([]*entity.User, error) {
	var users []*entity.User

	err := repoImpl.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	if users == nil {
		return []*entity.User{}, nil
	}

	return users, nil
}

func (repoImpl *UserRepositoryImpl) Delete(id string) error {
	return repoImpl.db.Delete(&entity.User{}, "id = ?", id).Error
}
