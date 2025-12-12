package builders

import (
	"errors"
	"go_template/internal/model"
	"go_template/internal/utils/apperror"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserOption func(*model.User) error

// UserBuilder construye una instancia de model.User con las opciones proporcionadas
// Valida la presencia de campos obligatorios y aplica valores predeterminados
func UserBuilder(opts ...UserOption) (*model.User, error) {
	u := &model.User{}

	for _, opt := range opts {
		if err := opt(u); err != nil {
			return nil, err
		}
	}

	// Validar campos NOT NULL obligatorios
	if u.Name == "" {
		return nil, apperror.FromCode(http.StatusBadRequest, "name is required")
	}
	if u.Email == "" {
		return nil, apperror.FromCode(http.StatusBadRequest, "email is required")
	}

	now := time.Now()
	u.ID = uuid.New()
	u.CreatedAt = &now
	u.UpdatedAt = &now

	return u, nil
}

// WithNameUser asigna el nombre del usuario validando su presencia
func WithNameUser(name string) UserOption {
	return func(u *model.User) error {
		if name == "" {
			return apperror.FromCode(http.StatusBadRequest, "Invalid product ID")
		}

		u.Name = name
		return nil
	}
}

// WithUserEmail asigna el email del usuario validando su formato y unicidad en la base de datos
func WithUserEmail(db *gorm.DB, email string) UserOption {
	return func(u *model.User) error {
		if email == "" {
			return errors.New("email inválido")
		}

		if res := db.First(&model.User{}, "email = ?", email); res.Error == nil {
			return apperror.FromCode(http.StatusBadRequest, "email already exists in the model")
		} else if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}

		u.Email = email
		return nil
	}
}
