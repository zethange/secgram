package chat

import "secgram/internal/models"

type Repository interface {
	GetAll() ([]*models.Chat, error)
	Get(id int64) (*models.Chat, error)
	GetAllByUserId(userId, page, limit uint64) ([]*models.Chat, error)
}
