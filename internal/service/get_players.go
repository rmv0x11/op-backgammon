package service

import (
	"github.com/rmv0x11/op-backgammon/internal/model"
)

func (s *Service) GetPlayers() ([]*model.Player, error) {
	return s.db.GetPlayers()
}
