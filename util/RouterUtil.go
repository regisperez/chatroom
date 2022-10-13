package util

import "github.com/gorilla/mux"

var(
	Router *mux.Router
)

func GetRouter () *mux.Router{
	return Router
}