package service

import "github.com/rmv0x11/op-backgammon/internal/storage"

func (s *Service) NewPlayer(player *storage.Player) (int64, error) {
	return s.db.NewPlayer(player)
}
