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
	DateTime time.Time `json:"datetime"`
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

func LastMessages(db *sql.DB) ([]Message, error) {
	rows, err := db.Query(
		"SELECT * FROM (SELECT * FROM chatroom.messages ORDER BY datetime desc limit 50) var1 ORDER BY id ASC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.User, &message.Message,&message.DateTime); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
