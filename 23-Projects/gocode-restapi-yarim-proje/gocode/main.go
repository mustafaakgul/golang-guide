package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/cihanozhan/internal/app"
	bookDB "github.com/cihanozhan/internal/book/db"
	userAPI "github.com/cihanozhan/internal/user/api"
	userDB "github.com/cihanozhan/internal/user/db"
	//"github.com/gorilla/handlers"
	//"github.com/joho.godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//BooksDB connection
	connBookDB, err := bookDB.CreateUsersDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	bDB := bookDB.NewDB(connBookDB)

	//UsersDB Connection
	connUserDB, err := userDB.CreateUsersDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	uDB := userDB.NewDB(connUserDB)
	//Application

	application := app.NewApplication()

	//Operations
	// bAPI := bookAPI.NewApp(bDB, application) //nesne ornegi olusturma
	// bAPI.GetBook(w,r)

	userAPI.NewApp(uDB, application)

	server := http.Server{
		Addr:    ":8080",
		Handler: application.Router,
	}

	log.Println("Listening on port", server.Addr)
	loggedRouter := handlers.LoggingHandler(os.Stdout, application.Router)
	http.ListenAndServe(server.Addr, loggedRouter)

}
