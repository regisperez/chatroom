package controller

import (
	"chatroom/consts"
	"chatroom/util"
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize()
	util.DatabaseTest = true
	SessionTest = true
	util.Router = a.Router
	ensureTableExists()
	code := m.Run()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := util.DB().Exec(consts.TableUsersCreationQuery); err != nil {
		log.Fatal(err)
	}
}