package chat

import "secgram/internal/models"

type UseCase interface {
	GetAllByUserId(chatId, limit, page uint64) ([]*models.Chat, error)
}
