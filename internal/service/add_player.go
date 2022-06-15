package service

import "github.com/rmv0x11/op-backgammon/internal/storage"

func (s *Service) AddPlayer(player *storage.Player) error {
	return s.db.InsertPlayer(player)
}
