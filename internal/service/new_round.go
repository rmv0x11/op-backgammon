package service

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

func (s *Service) NewRound(matchID, winnerID int64, isMars bool) (int64, error) {
	r := new(storage.Round)
	r.MatchID = sql.NullInt64{matchID, true}
	r.WinnerID = sql.NullInt64{winnerID, true}
	r.IsMars = sql.NullBool{isMars, true}
	return s.db.NewRound(r)
}
