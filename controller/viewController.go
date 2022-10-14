package controller

import (
	"net/http"
	"text/template"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	temp:= template.Must(template.ParseFiles("../view/login.html"))
	temp.Execute(w,r)
}

func ViewWelcome(w http.ResponseWriter, r *http.Request) {
	temp:= template.Must(template.ParseFiles("../view/welcome.html"))
	temp.Execute(w,r)
}