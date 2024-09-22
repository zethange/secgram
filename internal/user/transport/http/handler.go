package http

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"io"
	"net/http"
	u "secgram/internal/user"
	"secgram/internal/user/dto"
	"time"
)

type Handler struct {
	uc u.UseCase
}

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(jwt.MapClaims)

	username, ok := claims["username"].(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.uc.GetByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var data dto.LoginDTO
	if err = json.Unmarshal(body, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result, err := h.uc.Login(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//todo: add secure
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    result.Token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Path:     "/",
	})
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var data dto.RegisterDTO
	if err = json.Unmarshal(body, &data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registeredUser, err := h.uc.Register(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registeredUser)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	})
}

func NewHandler(uc u.UseCase) *Handler {
	return &Handler{uc: uc}
}
