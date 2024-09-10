package house

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

func (r *repository) CreateHouse(ctx context.Context, house *House) (*House, error) {
	var lastHouseId int
	query := "INSERT INTO house (house_name, adress) VALUES ($1, $2) RETURNING id "
	err := r.db.QueryRowContext(ctx, query, house.Adress, house.Name).Scan(&lastHouseId)
	if err != nil {
		return &House{}, err
	}
	house.ID = int64(lastHouseId)
	return house, nil
}
