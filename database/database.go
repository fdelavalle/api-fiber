package database

import (
	"log"

	"github.com/fdelavalle/api-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the database connection
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// ConnectDb connects to the database
func ConnectDb(connectionString string) {
	dsn := connectionString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to conenct to database \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Book{})

	DB = Dbinstance{
		Db: db,
	}

}
