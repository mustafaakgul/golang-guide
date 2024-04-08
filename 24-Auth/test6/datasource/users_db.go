package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test6/domain/users"
)

var (
	dbx *sql.DB
)

func init() {

	// username:password@tcp(host)/user_schema
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		"localhost",
		"5432",
		"postgres",
		"test6",
		"postgres",
	)

	dbx, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Errorf("cannot open database: %v", err)
	}

	//if !db.Migrator().HasTable(&Core{}) {
	log.Printf("Migrations started")
	dbx.AutoMigrate(
		&users.User{},
	)

	log.Println("database successfully configured")
}
