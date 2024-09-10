package message

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) *repository {
	return &repository{db: db}
}
func (r *repository) SaveMessage(c context.Context, message *Message) error {
	query := `INSERT INTO message(sender, receiver,content, timestamp) VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(c, query, message.Sender, message.Receiver, message.Content, message.TimeStamp)
	if err != nil {
		print("Error inserting message", err)
		return err
	}
	return nil
}
func (r *repository) GetMessageHistory(c context.Context, sender string, receiver string) ([]Message, error) {
	query := `SELECT sender, receiver, content, timestamp
				FROM message
				WHERE (sender = ? AND receiver = ?)
				OR (sender = ? AND receiver = ?)`
	rows, err := r.db.QueryContext(c, query, sender, receiver, receiver, sender)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Content, &msg.TimeStamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)

	}
	return messages, nil
}
