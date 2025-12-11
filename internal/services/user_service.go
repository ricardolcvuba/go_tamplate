package services

import (
	"go_template/internal/model"
	"go_template/internal/repository"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) UserService {
	return UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}
