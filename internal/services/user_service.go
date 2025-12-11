package services

import (
	"go_template/internal/dto"
	"go_template/internal/repository"

	"go_template/internal/utils"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) UserService {
	return UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return utils.ToUserDTOs(users), nil
}
