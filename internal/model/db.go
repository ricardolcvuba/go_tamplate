package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB establece una conexion a la base de datos SQLite con la URL proporcionada
func ConnectDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Printf("Failed to connect to database (%s): %v", databaseURL, err)
		return nil, err
	}

	sqlDB, _ := db.DB()
	if _, err := sqlDB.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		log.Printf("Failed to enable foreign keys: %v", err)
	}

	return db, nil
}

// ConnectDBTest establece una conexion a una base de datos SQLite en memoria para pruebas
func ConnectDBTest() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Printf("Failed to connect to test database: %v", err)
		return nil, err
	}

	sqlDB, _ := db.DB()
	if _, err := sqlDB.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		log.Printf("Failed to enable foreign keys: %v", err)
	}

	return db, nil
}

func MigrateSchemas(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}
