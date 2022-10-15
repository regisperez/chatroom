package util

import (
	"chatroom/common"
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	DatabaseTest bool
	ConnectionString string
)

func GetConnectionString(config string) string {
	var database common.Database
	if _, err := toml.DecodeFile(config, &database); err != nil {
		fmt.Println(err)
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", database.User, database.Password, database.Server, database.Port, database.Database)
	return connString
}

func DB() *sql.DB{

	if !DatabaseTest{
		ConnectionString = GetConnectionString("../config/database.toml")
	}else{
		ConnectionString = GetConnectionString("../config/database_test.toml")
	}

	var (
		err error
		DB *sql.DB
	)
	DB, err = sql.Open("mysql", ConnectionString +"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return DB
}