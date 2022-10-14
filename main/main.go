package main

import "chatroom/controller"

func main() {
	a := controller.App{}
	a.Initialize()
	a.Run("127.0.0.1")
}