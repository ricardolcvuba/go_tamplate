package utils

import (
	"go_template/internal/dto"
	"go_template/internal/model"
)

func ToUserDTOs(users []model.User) []dto.UserDTO {
	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}
	return userDTOs
}
