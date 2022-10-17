package controller

import (
	"net/http"
	"text/template"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	temp:= template.Must(template.ParseFiles("../view/login.html"))
	temp.Execute(w,r)
}

func ViewChatRoom(w http.ResponseWriter, r *http.Request) {
	if isInvalidSession(w,r){
		return
	}
	temp:= template.Must(template.ParseFiles("../view/chatroom.html"))
	temp.Execute(w,r)
}
func ViewAdmin(w http.ResponseWriter, r *http.Request) {
	if isInvalidSession(w,r){
		return
	}
	temp:= template.Must(template.ParseFiles("../view/admin.html"))
	temp.Execute(w,r)
}