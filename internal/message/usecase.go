package message

import "secgram/internal/models"

type UseCase interface {
	GetByChatId(userId, chatId, size, page uint64) ([]*models.Message, error)
	Create(message *models.Message) (*models.Message, error)
}
