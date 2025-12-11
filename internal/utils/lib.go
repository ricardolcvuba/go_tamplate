package utils

import (
	"go_template/internal/dto"
	"go_template/internal/model"
)

func ToUserDTO(user model.User) dto.UserDTO {
	return dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserDTOs(users []model.User) []dto.UserDTO {
	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = ToUserDTO(user)
	}
	return userDTOs
}
