package database

import (
	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/config"

	"log"

	"github.com/fabian-ss/Fiber-GORM-Swagger-Sqlite/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func ConnectDB() {

	db, err := gorm.Open(sqlite.Open("api_fiber.db"), &gorm.Config{})
	config.Errorhandle(err, "Error al conectar con la base de datos", 2)

	log.Println("Se establecio conección a la base de datos")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	//Migrations
	db.AutoMigrate(&models.User{}, &models.Nota{})

	Database = Dbinstance{Db: db}
}
