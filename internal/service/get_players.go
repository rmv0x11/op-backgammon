package service

import (
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

func (s *Service) GetPlayers() ([]*storage.Player, error) {
	return s.db.GetPlayers()
}
