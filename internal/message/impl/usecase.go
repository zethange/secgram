package impl

import (
	"secgram/internal/message"
	"secgram/internal/models"
	"slices"
)

type UseCaseImpl struct {
	repo message.Repository
}

func (u *UseCaseImpl) GetByChatId(userId, chatId, size, page uint64) ([]*models.Message, error) {
	users, err := u.repo.GetByChatId(userId, chatId, size, page)
	slices.Reverse(users)
	return users, err
}

func (u *UseCaseImpl) Create(message *models.Message) (*models.Message, error) {
	return u.repo.Create(message)
}

func NewUseCase(repo message.Repository) *UseCaseImpl {
	return &UseCaseImpl{repo: repo}
}
