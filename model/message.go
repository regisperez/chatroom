package model

import (
	"database/sql"
	"log"
	"time"
)

type Message struct {
	ID       int       `json:"id"`
	User     string    `json:"user"`
	Message  string    `json:"message"`
	DateTime time.Time `json:"date"`
}

func (message *Message) CreateMessage(db *sql.DB) error {
	insert, err := db.Exec("INSERT INTO messages (user, message, datetime) VALUES (?, ?, ?)", message.User, message.Message, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	id, err := insert.LastInsertId()

	message.ID = int(id)

	return nil
}
