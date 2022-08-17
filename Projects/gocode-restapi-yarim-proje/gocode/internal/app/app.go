package app

import (
	"fmt"
	"net/http"

	"github.com/cihanozhan/utils"
	"github.com/gorilla/mux"
)

type (
	Application struct {
		*mux.Router
	}
)

func NewApplication() *Application {
	appRouter := mux.NewRouter().StrictSlash(true)

	application := Application{
		Router: appRouter,
	}

	appRouter.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	appRouter.HandleFunc("/app", ApplicationHandler).Methods("GET")

	return &application
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(fmt.Errorf("Not Found"), http.StatusNotFound, w)
}

func ApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bu bir application router"))
}
