package entity

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key"`
	Name  string    `gorm:"size:500;not null"`
	Email string    `gorm:"size:500;not null;uniqueIndex"`
}

func NewUser(name string, email string) *User {
	return &User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
	}
}
