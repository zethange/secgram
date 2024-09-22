package server

import "secgram/internal/models"

const (
	Ping          = 0
	Pong          = 1
	NewMessage    = 2
	UpdateMessage = 3
	DeleteMessage = 4
	NewChat       = 5
	DeleteChat    = 6
	UpdateChat    = 7
	UserOnline    = 8
	UserOffline   = 9
)

type NewMessageStruct struct {
	Message *models.Message `json:"message"`
	ChatId  uint64          `json:"chat_id"`
}

type UserStatusStruct struct {
	UserId uint64 `json:"user_id"`
}

type MessageResponse struct {
	Type        uint64            `json:"type"`
	NewMessage  *NewMessageStruct `json:"new_message,omitempty"`
	UserOnline  *UserStatusStruct `json:"user_online,omitempty"`
	UserOffline *UserStatusStruct `json:"user_offline,omitempty"`
}
