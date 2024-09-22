package impl

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"secgram/internal/models"
	u "secgram/internal/user"
	"secgram/internal/user/dto"
	"secgram/pkg/util"
	"time"
)

type UseCaseImpl struct {
	repo u.Repository
}

func (u *UseCaseImpl) GetChatIncludeByUserId(userId uint64) ([]*models.User, error) {
	return u.repo.GetChatIncludeByUserId(userId)
}

func (u *UseCaseImpl) GetInChat(chatId uint64) ([]*models.User, error) {
	return u.repo.GetInChat(chatId)
}

func (u *UseCaseImpl) GetByUsername(username string) (*models.User, error) {
	return u.repo.GetByUsername(username)
}

func (u *UseCaseImpl) Register(registerDto *dto.RegisterDTO) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	registerDto.Password = string(hashedPassword)

	return u.repo.Create(registerDto)
}

func (u *UseCaseImpl) Get(userId uint64) (*models.User, error) {
	return u.repo.Get(userId)
}

func (u *UseCaseImpl) Login(loginDto *dto.LoginDTO) (*dto.LoginResponse, error) {
	user, err := u.repo.GetByUsername(loginDto.Username)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"id":       uint64(user.Id),
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(util.GetEnv("JWT_KEY", "pomodoro")))
	loginResponse := &dto.LoginResponse{Token: tokenString, User: user}
	return loginResponse, err
}

func (u *UseCaseImpl) JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(util.GetEnv("JWT_KEY", "pomodoro")), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user", token.Claims.(jwt.MapClaims)))

		next.ServeHTTP(w, r)
	}
}

func NewUseCase(repo u.Repository) *UseCaseImpl {
	return &UseCaseImpl{repo: repo}
}
