package models

import "time"

type User struct {
	Id        uint64    `json:"id"`
	FullName  string    `json:"full_name" db:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	LastSeen  time.Time `json:"last_seen" db:"last_seen"`
}
