package model

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (user *User) GetUser(db *sql.DB) error {
	return db.QueryRow("SELECT name, login, password FROM users WHERE id=?",
		user.ID).Scan(&user.Name, &user.Login,&user.Password)
}

func (user *User) UpdateUser(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE users SET name=?, login=?, password=? WHERE id=?",
			user.Name, user.Login, user.Password, user.ID)

	return err
}

func (user *User) DeleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=?", user.ID)

	return err
}

func (user *User) CreateUser(db *sql.DB) error {
	insert, err := db.Exec("INSERT INTO users (name, login, password) VALUES (?, ?, ?)", user.Name, user.Login, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := insert.LastInsertId()

	user.ID = int(id)

	return nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(
		"SELECT id, name, login, password FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Login,&user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (user *User) LoginUser(db *sql.DB) error {
	return db.QueryRow("SELECT name, login, password FROM users WHERE login=?",
		user.Login).Scan(&user.Name, &user.Login,&user.Password)
}

func HasAdminUser(db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE login=?","admin").Scan(&count)
	if err != nil {
		return 0
	}
	return count
}
