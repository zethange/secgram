package http

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"secgram/internal/chat"
	"secgram/internal/message"
	u "secgram/internal/user"
	"secgram/pkg/util"
)

type Handler struct {
	userUC    u.UseCase
	chatUC    chat.UseCase
	messageUC message.UseCase
}

func (h *Handler) GetMyChats(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("user").(jwt.MapClaims)["username"].(string)
	limit := uint64(util.ParseInt(r.URL.Query().Get("limit"), 10))
	page := uint64(util.ParseInt(r.URL.Query().Get("page"), 1))

	user, err := h.userUC.GetByUsername(username)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	chats, err := h.chatUC.GetAllByUserId(user.Id, limit, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(chats)
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("user").(jwt.MapClaims)["username"].(string)
	chatId := uint64(util.ParseInt(r.URL.Query().Get("chatId"), 1))
	limit := uint64(util.ParseInt(r.URL.Query().Get("limit"), 10))
	page := uint64(util.ParseInt(r.URL.Query().Get("page"), 1))

	user, err := h.userUC.GetByUsername(username)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	messages, err := h.messageUC.GetByChatId(chatId, user.Id, limit, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}

func NewHandler(userUC u.UseCase, chatUC chat.UseCase, messageUC message.UseCase) *Handler {
	return &Handler{userUC: userUC, chatUC: chatUC, messageUC: messageUC}
}
