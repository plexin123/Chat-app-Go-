package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	//Connection to MySQL
	dsn := "postgres://user:password@localhost:5432/chat2_postgres?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err

	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
