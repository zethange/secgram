package client

const (
	Ping          = 0
	Pong          = 1
	NewMessage    = 2
	UpdateMessage = 3
	DeleteMessage = 4
	NewChat       = 5
	DeleteChat    = 6
	UpdateChat    = 7
)

type MessageDto struct {
	Type       uint64 `json:"type"`
	NewMessage *struct {
		Message string `json:"message"`
		ChatId  uint64 `json:"chat_id"`
	} `json:"new_message"`
}
