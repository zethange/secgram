package models

import "time"

type Chat struct {
	Id              uint64     `json:"id"`
	Type            string     `json:"type"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	Messages        []*Message `json:"messages"`
	Name            string     `json:"name"`
	LastMessageDate time.Time  `json:"last_message_date" db:"last_message_date"`
	Members         []*User    `json:"members"`
}
