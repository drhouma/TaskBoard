package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	maxOpenConn     = 50
	maxIdleConn     = 25
	connMaxLifetime = time.Duration(30)
)

type DataBase struct {
	DB *sqlx.DB
}

func New(dsn string) (*DataBase, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database connection: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(connMaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return &DataBase{
		DB: db,
	}, nil
}

func (db *DataBase) Close() error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("close database: %w", err)
	}
	return nil
}
