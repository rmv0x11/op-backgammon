package service

import "github.com/rmv0x11/op-backgammon/internal/storage"

type Service struct {
	db *storage.Database
}

func NewService() *Service {
	dsn := "backgammon.db" //TODO read from config

	db := storage.NewStorage(dsn)
	return &Service{db}
}
