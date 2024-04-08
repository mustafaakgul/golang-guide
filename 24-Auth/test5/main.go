package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//get it from .env file
var SECRET = []byte("super-secret-auth-key")
var api_key = "1234"

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized" + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	})
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "super secret area")
}

func main() {
	http.HandleFunc("/jwt", GetJWT)
	http.Handle("/api", ValidateJWT(Home))

	http.ListenAndServe(":3500", nil)
}
