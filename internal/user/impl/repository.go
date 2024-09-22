package impl

import (
	"github.com/jmoiron/sqlx"
	"secgram/internal/chat"
	"secgram/internal/models"
	"secgram/internal/user/dto"
)

type PostgresRepository struct {
	db *sqlx.DB
	chat.Repository
}

func (r *PostgresRepository) GetChatIncludeByUserId(userId uint64) ([]*models.User, error) {
	var users []*models.User
	err := r.db.Select(&users, `
		SELECT u.*
		FROM users u
		JOIN chats_users cu ON u.id = cu.user_id
		JOIN chats c ON cu.chat_id = c.id
		WHERE c.id IN (
				SELECT cu2.chat_id
				FROM chats_users cu2
				WHERE cu2.user_id = $1
		)
		AND u.id != $1;
	`, userId)
	return users, err
}

func (r *PostgresRepository) GetInChat(chatId uint64) ([]*models.User, error) {
	var users []*models.User
	err := r.db.Select(&users, `
		SELECT u.*
		FROM users u
		JOIN public.chats_users cu on u.id = cu.user_id
		WHERE cu.chat_id = $1
	`, chatId)
	return users, err
}

func (r *PostgresRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	return &user, err
}

func (r *PostgresRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Select(&users, "SELECT * FROM users")
	return users, err
}

func (r *PostgresRepository) Get(id uint64) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return &user, err
}

func (r *PostgresRepository) Create(user *dto.RegisterDTO) (*models.User, error) {
	var savedUser models.User
	err := r.db.Get(&savedUser, "INSERT INTO users (full_name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *", user.FullName, user.Username, user.Email, user.Password)
	return &savedUser, err
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
