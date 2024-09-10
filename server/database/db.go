package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	//Connection to MySQL
	dsn := "root:password@tcp(localhost:3307)/chat2-mysql?tls=false"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
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
