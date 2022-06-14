package app

import (
	"context"
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

type Implementation struct {
	db *storage.Database
}

func NewBackgammonAPI(ctx context.Context) *Implementation {
	dsn := "./backgammon.db"
	db := storage.NewStorage(getSQLite(dsn))
	return &Implementation{db: db}
}

func getSQLite(dsn string) *sql.DB {
	sqliteDB, _ := sql.Open("sqlite3", dsn)
	return sqliteDB
}

func (i *Implementation) Close() error {
	return i.db.Close()
}
