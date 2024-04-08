package models

import (
	"fmt"
	"github.com/joho/godotenv"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DBInstance *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		"localhost",
		"5432",
		"postgres",
		"test8",
		"postgres",
	)

	DBInstance, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Errorf("cannot open database: %v", err)
	}

	DBInstance.AutoMigrate(&User{})
}
