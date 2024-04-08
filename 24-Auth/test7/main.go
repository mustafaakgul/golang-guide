package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

var DB *gorm.DB

func main() {

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		"localhost",
		"5432",
		"postgres",
		"test6",
		"postgres",
	)

	DB, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		fmt.Errorf("cannot open database: %v", err)
	}

	DB.AutoMigrate(&User{})

	/*usertest := User{Email: "hello3@hello.com", Password: "hello"}
	DB.Create(&usertest)*/

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.POST("/signup", Signup)
	r.Run()
}

func Signup(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	usertest2 := User{Email: "hello2@hello.com", Password: "hello"}
	DB.Create(&usertest2)

	user := User{Email: body.Email, Password: string(hash)}
	result := DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
