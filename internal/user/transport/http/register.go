package http

import (
	"net/http"
	u "secgram/internal/user"
)

func RegisterHTTPMethods(uc u.UseCase, mux *http.ServeMux) {
	h := NewHandler(uc)

	mux.HandleFunc("GET /api/users/me", uc.JWTMiddleware(h.GetMe))

	mux.HandleFunc("POST /api/auth/login", h.Login)
	mux.HandleFunc("POST /api/auth/register", h.Register)
	mux.HandleFunc("POST /api/auth/logout", h.Logout)

	//mux.HandleFunc("")
}
