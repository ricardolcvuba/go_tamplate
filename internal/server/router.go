package server

import (
	"go_template/internal/handlers"
	"go_template/internal/middleware"
	"go_template/internal/repository"
	"go_template/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	middleware.SetupMiddlewares(r)

	userRepo := repository.NewUserRepository(db)

	userService := services.NewUserService(&userRepo)

	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users", userHandler.GetAllUsers)

	return r
}
