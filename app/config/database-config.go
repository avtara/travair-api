package config

import (
	"fmt"
	_reservRepo "github.com/avtara/travair-api/repository/databases/reservations"
	_unitsRepo "github.com/avtara/travair-api/repository/databases/units"
	_usersRepo "github.com/avtara/travair-api/repository/databases/users"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}


	var dbName string
	if os.Getenv("ENV") == "TESTING"{
		dbName = os.Getenv("DB_NAME_TESTING")
	} else {
		dbName = os.Getenv("DB_NAME")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error())
	}
	dbMigrate(db)

	return db
}

func dbMigrate(db *gorm.DB) {

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&_usersRepo.Users{}, &_unitsRepo.Units{}, &_unitsRepo.Photos{}, &_reservRepo.Reservation{})
}
