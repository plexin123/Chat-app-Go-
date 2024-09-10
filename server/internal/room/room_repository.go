package room

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
	return &repository{
		db: db,
	}
}

// CREATE TABLE Room (
//     room_id INT AUTO_INCREMENT PRIMARY KEY,
//     room_name VARCHAR(50) NOT NULL,
//     house_id INT NOT NULL,
//     user_id INT,
//     status BOOLEAN NOT NULL,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     FOREIGN KEY (house_id) REFERENCES House(house_id),
//     FOREIGN KEY (user_id) REFERENCES User(user_id)
// );

func (rep *repository) CreateRoom(ctx context.Context, room *Room) (*Room, error) {
	var lastRoomId int

	query := `INSERT INTO room (room_name,house_id, user_id,status)
			VALUES (? , ? , ? , ? ) RETURNING ID `

	err := rep.db.QueryRowContext(ctx, query, room.Name, room.House_id, room.User_id, room.Status).Scan(&lastRoomId)
	if err != nil {
		return &Room{}, err
	}
	room.ID = int64(lastRoomId)
	return room, nil
}
