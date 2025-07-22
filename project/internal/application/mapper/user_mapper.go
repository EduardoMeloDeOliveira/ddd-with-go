package mapper

import (
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/application/dto"
	"github.com/EduardoMeloDeOliveira/ddd-with-go/project/internal/domain/entity"
)

func ToUserResponse(user *entity.User) *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}
