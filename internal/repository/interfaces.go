package repository

import "go_template/internal/model"

type UserRepositoryInterface interface {
	GetAll() ([]model.User, error)
}
