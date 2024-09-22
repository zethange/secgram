package http

import (
	"net/http"
	"secgram/internal/chat"
	"secgram/internal/message"
	u "secgram/internal/user"
)

func RegisterHTTPMethods(userUC u.UseCase, chatUC chat.UseCase, messageUC message.UseCase, mux *http.ServeMux) {
	h := NewHandler(userUC, chatUC, messageUC)

	mux.HandleFunc("GET /api/chats", userUC.JWTMiddleware(h.GetMyChats))
	mux.HandleFunc("GET /api/chats/messages", userUC.JWTMiddleware(h.GetMessages))

	//mux.HandleFunc("")
}
