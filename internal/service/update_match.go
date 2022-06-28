package service

import (
	"github.com/rmv0x11/op-backgammon/internal/mappers"
	"github.com/rmv0x11/op-backgammon/internal/model"
)

func (s *Service) UpdateMatch(m *model.Match) error {
	if m.Length <= m.PlayerOnePoints || m.Length <= m.PlayerTwoPoints {
		m.Status = "The match is over."
		//update elo first player

		//update elo two player

	} else {
		m.Status = "Match in progress."
	}
	return s.db.UpdateMatch(mappers.MatchesForDB(m))
}
