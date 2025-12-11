package services

import "go_template/internal/repository"

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) UserService {
	return UserService{repo: repo}
}
