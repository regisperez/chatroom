package main

import "chatroom/controller"

func main() {
	a := controller.App{}
	a.Initialize()
	a.Run()
}