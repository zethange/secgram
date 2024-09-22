package websocket

import (
	"context"
	"encoding/json"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"secgram/internal/chat"
	client2 "secgram/internal/chat/dto/client"
	"secgram/internal/chat/dto/server"
	"secgram/internal/message"
	"secgram/internal/models"
	u "secgram/internal/user"
	"secgram/pkg/util"
	"sync"
)

type handler struct {
	uc        chat.UseCase
	userUC    u.UseCase
	messageUC message.UseCase

	clients     map[*websocket.Conn]*models.User
	clientsConn map[uint64]*websocket.Conn

	broadcast  chan *server.MessageResponse
	register   chan *User
	unregister chan *websocket.Conn
	mutex      sync.Mutex
}

type User struct {
	conn *websocket.Conn
	user *models.User
}

func (h *handler) ws(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{})
	if err != nil {
		log.Println(err)
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
		return
	}

	client := conn

	claims := r.Context().Value("user").(jwt.MapClaims)
	username := claims["username"].(string)

	user, err := h.userUC.GetByUsername(username)
	if err != nil {
		log.Println(err)
		return
	}

	h.register <- &User{
		conn: client,
		user: user,
	}

	defer func() {
		h.unregister <- client
	}()

	log.Printf("user [%s](id: %v) connected to ws gateway\n", username, user.Id)

	for {
		ctx, cancel := context.WithCancel(context.Background() /*, time.Second*10*/)

		var msg client2.MessageDto
		err := wsjson.Read(context.Background(), client, &msg)
		if err != nil {
			log.Println(err)
			cancel()
			break
		}

		switch msg.Type {
		case client2.NewMessage:
			usersInChat, err := h.userUC.GetInChat(msg.NewMessage.ChatId)
			if err != nil {
				continue
			}

			if !util.Some(usersInChat, func(u *models.User) bool {
				return user.Id == u.Id
			}) {
				continue
			}

			savedMessage, err := h.messageUC.Create(&models.Message{
				Content: msg.NewMessage.Message,
				UserId:  user.Id,
				ChatId:  msg.NewMessage.ChatId,
			})
			if err != nil {
				log.Println(err)
			}

			for _, u := range usersInChat {
				if c, ok := h.clientsConn[u.Id]; ok {
					response := server.MessageResponse{
						Type:       server.NewMessage,
						NewMessage: &server.NewMessageStruct{Message: savedMessage, ChatId: msg.NewMessage.ChatId},
					}
					wsjson.Write(ctx, c, response)
				}
			}
		}

		cancel()
	}
}

func newHandler(uc chat.UseCase, userUC u.UseCase, messageUC message.UseCase) *handler {
	h := &handler{
		uc:          uc,
		userUC:      userUC,
		messageUC:   messageUC,
		clients:     make(map[*websocket.Conn]*models.User),
		clientsConn: make(map[uint64]*websocket.Conn),
		broadcast:   make(chan *server.MessageResponse),
		register:    make(chan *User),
		unregister:  make(chan *websocket.Conn),
	}

	go func() {
		log.Println("init ws event loop")
		for {
			select {
			case client := <-h.register:
				go func() {
					users, err := h.userUC.GetChatIncludeByUserId(client.user.Id)
					if err != nil {
						log.Println(err)
						return
					}

					h.mutex.Lock()
					for _, user := range users {
						if c, ok := h.clientsConn[user.Id]; ok {
							wsjson.Write(context.Background(), c, &server.MessageResponse{
								Type: server.UserOnline,
								UserOnline: &server.UserStatusStruct{
									UserId: client.user.Id,
								},
							})
						}
					}
					h.mutex.Unlock()
				}()

				h.mutex.Lock()
				h.clients[client.conn] = client.user
				h.clientsConn[client.user.Id] = client.conn
				h.mutex.Unlock()
			case client := <-h.unregister:
				user := *h.clients[client]
				go func() {
					users, err := h.userUC.GetChatIncludeByUserId(user.Id)
					if err != nil {
						return
					}

					h.mutex.Lock()
					for _, u := range users {
						if c, ok := h.clientsConn[u.Id]; ok {
							wsjson.Write(context.Background(), c, &server.MessageResponse{
								Type: server.UserOffline,
								UserOffline: &server.UserStatusStruct{
									UserId: user.Id,
								},
							})
						}
					}
					h.mutex.Unlock()
				}()

				h.mutex.Lock()
				if user, ok := h.clients[client]; ok {
					if _, ok := h.clientsConn[user.Id]; ok {
						delete(h.clientsConn, user.Id)
					}
					delete(h.clients, client)
				}
				h.mutex.Unlock()
			case msg := <-h.broadcast:
				data, _ := json.Marshal(msg)

				h.mutex.Lock()
				for client := range h.clients {
					_ = client.Write(context.Background(), websocket.MessageText, data)
				}
				h.mutex.Unlock()
			}
		}
	}()
	return h
}
