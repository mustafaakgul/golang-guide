package api

import (
	"net/http"

	app "github.com/cihanozhan/internal/app"
	"github.com/cihanozhan/internal/model"
	"github.com/cihanozhan/internal/user/db"
	"github.com/gorilla/mux"
)

type (
	App struct {
		*mux.Router
		db db.UsersDB
	}
)

func NewApp(db db.UsersDB, application *app.Application) *App {
	app := App{
		Router: application.Router,
		db:     db,
	}
	app.Router.HandleFunc("/health", app.HealtCheck).Methods("GET")

	for _, r := range app.GetRoutes() {
		app.Router.HandleFunc(r.Pattern, r.HandlerFunc).Methods(r.Method)
	}

	return &app
}

func (app *App) GetRoutes() []model.Route {
	return []model.Route{
		{
			Method:      "GET",
			Pattern:     "/users",
			HandlerFunc: app.GetUsers,
		},
		{
			Method:      "POST",
			Pattern:     "/users",
			HandlerFunc: app.AddUser,
		},
	}
}

func (app *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users..."))
}

func (app *App) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Added a user"))
}

func (app *App) HealtCheck(w http.ResponseWriter, r *http.Request) {
	if err := app.db.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}
	w.WriteHeader(http.StatusOK) // 200
}
