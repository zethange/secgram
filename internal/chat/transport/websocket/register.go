package websocket

import (
	"net/http"
	"secgram/internal/chat"
	"secgram/internal/message"
	u "secgram/internal/user"
)

func RegisterHTTPMethods(chatUC chat.UseCase, userUC u.UseCase, messageUC message.UseCase, mux *http.ServeMux) {
	h := newHandler(chatUC, userUC, messageUC)
	mux.HandleFunc("/api/ws", userUC.JWTMiddleware(h.ws))
}
