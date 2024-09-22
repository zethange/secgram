package dto

import "secgram/internal/models"

type LoginDTO struct {
	// username or email, for default username
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}
