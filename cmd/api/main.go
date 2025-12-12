package main

import (
	"go_template/internal/config"
	"go_template/internal/model"
	"go_template/internal/server"
	"log"
)

func main() {
	config := config.GetConfig()
	db, err := model.ConnectDB(config.DatabaseURL)

	model.MigrateSchemas(db)

	if err != nil {
		log.Println("Connection refused to database")
		panic(err)
	}

	r := server.NewRouter(db)
	r.Run(":" + config.Port)
}
