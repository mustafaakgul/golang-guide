package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

//var mySigningKey = os.Get("JWT_SECRET")
var mySigningKey = []byte("secret")

func GenerateJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "mustafa"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, _ := token.SignedString(mySigningKey)
	return tokenString
}

func main() {
	token := GenerateJWT()

	print(token)
}
