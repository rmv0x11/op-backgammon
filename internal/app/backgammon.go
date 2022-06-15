package app

import (
	"context"
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"log"
	"os"
)

type Implementation struct {
	db *storage.Database
}

func NewBackgammonAPI(ctx context.Context) *Implementation {
	dsn := "backgammon.db"

	//if _, err := os.Open(dsn); err != nil {
	//	err = createDBFile(dsn)
	//	if err != nil {
	//		return nil
	//	}
	//}
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

func createDBFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()

	return err
}
