package impl

import (
	"secgram/internal/chat"
	"secgram/internal/models"
)

type UseCaseImpl struct {
	repo chat.Repository
}

func (u *UseCaseImpl) GetAllByUserId(chatId, limit, page uint64) ([]*models.Chat, error) {
	return u.repo.GetAllByUserId(chatId, page, limit)
}

func NewUseCase(repo chat.Repository) *UseCaseImpl {
	return &UseCaseImpl{repo: repo}
}
