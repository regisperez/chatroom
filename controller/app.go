package controller

import (
	"chatroom/consts"
	"chatroom/model"
	"chatroom/util"
	"chatroom/websocket"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var messages []string

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

	pool := websocket.NewPool()
	go pool.Start()

	a.Router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}



func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	c, err := r.Cookie("user_name")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return
		}
		// For any other type of error, return a bad request status
		return
	}
	userName := c.Value

	client := &websocket.Client{ID: userName,
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
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

