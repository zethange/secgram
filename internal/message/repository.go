package message

import "secgram/internal/models"

type Repository interface {
	GetByChatId(userId, chatId, limit, page uint64) ([]*models.Message, error)
	Create(message *models.Message) (*models.Message, error)
}
