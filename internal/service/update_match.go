package service

import "github.com/rmv0x11/op-backgammon/internal/model"

func (s *Service) UpdateMatch(m *model.Match) error {

	return s.db.UpdateMatch(m)
}
