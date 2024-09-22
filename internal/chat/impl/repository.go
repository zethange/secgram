package impl

import (
	"github.com/jmoiron/sqlx"
	"secgram/internal/chat"
	"secgram/internal/models"
)

type PostgresRepository struct {
	db *sqlx.DB
	chat.Repository
}

func (r *PostgresRepository) GetAll() ([]*models.Chat, error) {
	var chats []*models.Chat
	err := r.db.Select(&chats, "SELECT * FROM chats")
	return chats, err
}

func (r *PostgresRepository) Get(id int64) (*models.Chat, error) {
	var c models.Chat
	err := r.db.Get(&c, "SELECT * FROM chats WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &c, err
}

func (r *PostgresRepository) GetAllByUserId(userId, page, limit uint64) ([]*models.Chat, error) {
	var chats []*models.Chat

	offset := (page - 1) * limit
	err := r.db.Select(&chats, `
		SELECT c.*,
    COALESCE((SELECT u.full_name
     FROM users u
              JOIN chats_users cu ON u.id = cu.user_id
     WHERE cu.chat_id = c.id AND cu.user_id <> $1
     LIMIT 1), 'Saved messages') AS name,
		COALESCE((SELECT MAX(m.created_at) FROM messages m WHERE m.chat_id = c.id), now()) as last_message_date
		FROM chats c
    JOIN chats_users cu ON c.id = cu.chat_id
		WHERE cu.user_id = $1
		LIMIT $2 OFFSET $3
	`, userId, limit, offset)

	// TODO избавится от N+1
	for _, c := range chats {
		err := r.db.Select(&c.Members, `
			SELECT u.*
			FROM users u
			JOIN public.chats_users cu on u.id = cu.user_id
			WHERE cu.chat_id = $1
			AND u.id != $2
		`, c.Id, userId)
		if err != nil {
			return chats, err
		}
	}

	return chats, err
}

func (r *PostgresRepository) Create(chat *models.Chat) error {
	return nil
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
