package models

import "time"

type Message struct {
	Id           uint64    `json:"id"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UserId       uint64    `json:"user_id" db:"user_id"`
	UserFullName string    `json:"user_full_name" db:"user_full_name"`
	ChatId       uint64    `json:"chat_id" db:"chat_id"`
}
