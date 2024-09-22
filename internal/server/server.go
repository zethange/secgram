package server

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"secgram/pkg/util"

	userimpl "secgram/internal/user/impl"
	userhttp "secgram/internal/user/transport/http"

	chatimpl "secgram/internal/chat/impl"
	chathttp "secgram/internal/chat/transport/http"
	chatws "secgram/internal/chat/transport/websocket"

	messageimpl "secgram/internal/message/impl"
)

type Server struct {
	db *sqlx.DB
}

func NewServer() *Server {
	return &Server{}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://secgram.cubteam.ru")                    // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE, OPTIONS") // Allowed methods
		w.Header().Set("Access-Control-Allow-Credentials", "true")                                     //
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")                                 // Allowed headers

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) initDb() error {
	databaseUrl := util.GetEnv("DATABASE_URL", "postgres://zethange:tomato@localhost:5432/secgram?sslmode=disable")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	conn, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		return err
	}
	s.db = conn
	return nil
}

func (s *Server) Run(port string) error {
	if err := s.initDb(); err != nil {
		return err
	} else {
		log.Println("connection to db established")
	}

	mux := http.NewServeMux()

	messageUC := messageimpl.NewUseCase(messageimpl.NewPostgresRepository(s.db))

	userUC := userimpl.NewUseCase(userimpl.NewPostgresRepository(s.db))
	userhttp.RegisterHTTPMethods(userUC, mux)

	chatUC := chatimpl.NewUseCase(chatimpl.NewPostgresRepository(s.db))
	chatws.RegisterHTTPMethods(chatUC, userUC, messageUC, mux)
	chathttp.RegisterHTTPMethods(userUC, chatUC, messageUC, mux)

	log.Println("try to starting server on " + port)
	return http.ListenAndServe(port, corsMiddleware(mux))
}
