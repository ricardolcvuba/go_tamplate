package shared

import (
	"go_template/internal/model"
	"go_template/internal/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTest() (*gin.Engine, *gorm.DB, error) {
	db, err := model.ConnectDBTest()
	if err != nil {
		return nil, nil, err
	}

	router := server.NewRouter(db)

	model.MigrateSchemas(db)

	return router, db, nil
}
