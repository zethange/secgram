package chat

import (
	"secgram/internal/models"
	"secgram/internal/user/dto"
)

type Repository interface {
	GetAll() ([]*models.User, error)
	Get(id uint64) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Create(user *dto.RegisterDTO) (*models.User, error)
	GetInChat(chatId uint64) ([]*models.User, error)
	GetChatIncludeByUserId(userId uint64) ([]*models.User, error)
}
