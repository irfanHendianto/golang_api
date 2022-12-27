package config

import (
	"fmt"
	"os"

	"github.com/golang_api/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetUpDatabaseConnection is creating a new connection to our Database
func SetUpDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	fmt.Print(errEnv)
	if errEnv != nil {
		panic("Failed to load Database")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entity.User{})
	return db
}

// CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
