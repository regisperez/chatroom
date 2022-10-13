package main

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

//func clearTable() {
//	util.DB().Exec("DELETE FROM users")
//	util.DB().Exec("ALTER TABLE users AUTO_INCREMENT = 1;")
//}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
  id SERIAL,
  name TEXT NOT NULL,
  login TEXT NOT NULL,
  password TEXT NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)`

//func TestEmptyTable(t *testing.T) {
//	clearTable()
//
//	req, _ := http.NewRequest("GET", "/users", nil)
//	response := executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	if body := response.Body.String(); body != "[]" {
//		t.Errorf("Expected an empty array. Got %s", body)
//	}
//}
//
//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//	rr := httptest.NewRecorder()
//	a.Router.ServeHTTP(rr, req)
//
//	return rr
//}
//
//func checkResponseCode(t *testing.T, expected, actual int) {
//	if expected != actual {
//		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//	}
//}
//
//func TestGetNonExistentUser(t *testing.T) {
//	clearTable()
//
//	req, _ := http.NewRequest("GET", "/user/11", nil)
//	response := executeRequest(req)
//
//	checkResponseCode(t, http.StatusNotFound, response.Code)
//
//	var m map[string]string
//	json.Unmarshal(response.Body.Bytes(), &m)
//	if m["error"] != "User not found" {
//		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
//	}
//}
//
//func TestCreateUser(t *testing.T) {
//	clearTable()
//
//	var jsonStr = []byte(`{"name":"test user", "login": "test", "password":"pass"}`)
//	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
//	req.Header.Set("Content-Type", "application/json")
//
//	response := executeRequest(req)
//	checkResponseCode(t, http.StatusCreated, response.Code)
//
//	var m map[string]interface{}
//	json.Unmarshal(response.Body.Bytes(), &m)
//
//	if m["name"] != "test user" {
//		t.Errorf("Expected user name to be 'test user'. Got '%v'", m["name"])
//	}
//
//	if m["login"] != "test" {
//		t.Errorf("Expected user login to be 'test'. Got '%v'", m["login"])
//	}
//
//	if m["password"] != "pass" {
//		t.Errorf("Expected user login to be 'pass'. Got '%v'", m["password"])
//	}
//
//	if m["id"] != 1.0 {
//		t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
//	}
//}
//
//func TestGetUser(t *testing.T) {
//	clearTable()
//	addUsers(1)
//
//	req, _ := http.NewRequest("GET", "/user/1", nil)
//	response := executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//}
//
//func addUsers(count int) {
//	if count < 1 {
//		count = 1
//	}
//
//	for i := 0; i < count; i++ {
//		util.DB().Exec("INSERT INTO users(name, login,password) VALUES(?, ?, ?)", "User "+strconv.Itoa(i), "Login "+strconv.Itoa(i),"Password "+strconv.Itoa(i))
//	}
//}
//
//func TestUpdateUser(t *testing.T) {
//
//	clearTable()
//	addUsers(1)
//
//	req, _ := http.NewRequest("GET", "/user/1", nil)
//	response := executeRequest(req)
//	var originalUser map[string]interface{}
//	json.Unmarshal(response.Body.Bytes(), &originalUser)
//
//	var jsonStr = []byte(`{"name":"test user - updated name", "login": "test updated","password": "pass updated"}`)
//	req, _ = http.NewRequest("PUT", "/user/1", bytes.NewBuffer(jsonStr))
//	req.Header.Set("Content-Type", "application/json")
//
//	response = executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	var m map[string]interface{}
//	json.Unmarshal(response.Body.Bytes(), &m)
//
//	if m["id"] != originalUser["id"] {
//		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser["id"], m["id"])
//	}
//
//	if m["name"] == originalUser["name"] {
//		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser["name"], m["name"], m["name"])
//	}
//
//	if m["login"] == originalUser["login"] {
//		t.Errorf("Expected the login to change from '%v' to '%v'. Got '%v'", originalUser["login"], m["login"], m["login"])
//	}
//}
//
//func TestDeleteUser(t *testing.T) {
//	clearTable()
//	addUsers(1)
//
//	req, _ := http.NewRequest("GET", "/user/1", nil)
//	response := executeRequest(req)
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	req, _ = http.NewRequest("DELETE", "/user/1", nil)
//	response = executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//
//	req, _ = http.NewRequest("GET", "/user/1", nil)
//	response = executeRequest(req)
//	checkResponseCode(t, http.StatusNotFound, response.Code)
//}