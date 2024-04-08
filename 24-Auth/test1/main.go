package main

import (
	"gorm.io/driver/sqlite"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
}

func insertUser(username string, password string, age uint) (*User, error) {
	user := User{
		Username: username,
		Password: password,
		Age:      age,
	}
	if res := DB.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
func findUserByUsername(username string) (*User, error) {
	var user User
	if res := DB.Where("username = ?", username).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
func findUserByID(id uint) (*User, error) {
	var user User
	if res := DB.Find(&user, id); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	DB = db
	_, _ = insertUser("david", "123", 30)
	_, _ = insertUser("sophia", "321", 25)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/login", login)

	r.Use(verifyToken)
	r.GET("/info", getUserInfo)

	r.GET("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome to admin dashboard!")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "love",
		"user2": "god",
		"user3": "sex",
	}))
	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func basicAuth(c *gin.Context) {
	// Get the Basic Authentication credentials
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "testuser" && password == "testpass" {
		print("user authenticated")
	} else {
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}

func verifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	id, username, err := validateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	c.Set("id", id)
	c.Set("username", username)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}
func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
func getSession(c *gin.Context) (uint, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return 0, "", false
	}
	username, ok := c.Get("username")
	if !ok {
		return 0, "", false
	}
	return id.(uint), username.(string), true
}
