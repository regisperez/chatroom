package controller

import (
	"chatroom/consts"
	"chatroom/model"
	"chatroom/util"
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
	ensureTablesExists()
	ensureHasAdminUser()
	a.initializeRoutes()
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/", ViewLogin).Methods("GET")
	a.Router.HandleFunc("/chatroom", ViewChatRoom).Methods("GET")
	a.Router.HandleFunc("/users", GetUsers).Methods("GET")
	a.Router.HandleFunc("/user", CreateUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", GetUser).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id:[0-9]+}", DeleteUser).Methods("DELETE")
	a.Router.HandleFunc("/user/login", LoginUser).Methods("POST")
	a.Router.HandleFunc("/user/logout", LogoutUser).Methods("POST")

	a.Router.HandleFunc("/message", CreateMessage).Methods("POST")
	a.Router.HandleFunc("/lastMessages", LastMessages).Methods("GET")
}

func ensureTablesExists() {
	if _, err := util.DB().Exec(consts.TableUsersCreationQuery); err != nil {
		log.Fatal(err)
	}

	if _, err := util.DB().Exec(consts.TableMessagesCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func ensureHasAdminUser(){
	if model.HasAdminUser(util.DB()) == 0{
		if _, err := util.DB().Exec("INSERT INTO users(name, login,password) VALUES(?, ?, ?)", "Admin", "admin","$2a$08$tdBCe0L6QuocnBINJ7XZmODa4GdTNmp2qtsBqVqCbYoIxD.PBGFfW");err != nil {
			log.Fatal(err)
		}
	}
}

