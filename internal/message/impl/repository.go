package impl

import (
	"github.com/jmoiron/sqlx"
	"secgram/internal/models"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func (r *PostgresRepository) Create(message *models.Message) (*models.Message, error) {
	var msg models.Message
	err := r.db.Get(&msg, `
		WITH inserted AS (
			INSERT INTO messages (content, user_id, chat_id)
			VALUES ($1, $2, $3)
			RETURNING *
		)
		SELECT *, (SELECT full_name FROM users WHERE id = inserted.user_id) as user_full_name
		FROM inserted
	`, message.Content, message.UserId, message.ChatId)
	return &msg, err
}

func (r *PostgresRepository) GetByChatId(chatId, userId, limit, page uint64) ([]*models.Message, error) {
	var messages []*models.Message

	offset := (page - 1) * limit
	err := r.db.Select(&messages, `
			SELECT m.*, u.full_name as user_full_name
			FROM messages m 
			JOIN public.users u on u.id = m.user_id
			JOIN public.chats c on c.id = m.chat_id
			JOIN public.chats_users cu on c.id = cu.chat_id
			WHERE m.chat_id = $1
			AND cu.user_id = $4
			ORDER BY m.created_at DESC
			LIMIT $2 OFFSET $3
		`, chatId, limit, offset, userId)

	return messages, err
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}
