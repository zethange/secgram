package chat

import (
	"net/http"
	"secgram/internal/models"
	"secgram/internal/user/dto"
)

type UseCase interface {
	Get(userId uint64) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Register(registerDto *dto.RegisterDTO) (*models.User, error)
	Login(loginDto *dto.LoginDTO) (*dto.LoginResponse, error)
	GetInChat(chatId uint64) ([]*models.User, error)
	GetChatIncludeByUserId(userId uint64) ([]*models.User, error)
	JWTMiddleware(next http.HandlerFunc) http.HandlerFunc
}
