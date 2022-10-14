package controller

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/users", GetUsers).Methods("GET")
	a.Router.HandleFunc("/user", CreateUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", GetUser).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id:[0-9]+}", DeleteUser).Methods("DELETE")
	a.Router.HandleFunc("/user/login", LoginUser).Methods("POST")
}
