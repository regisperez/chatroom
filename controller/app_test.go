package controller

import (
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
	if _, err := util.DB().Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
  id SERIAL,
  name TEXT NOT NULL,
  login TEXT NOT NULL,
  password TEXT NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)`