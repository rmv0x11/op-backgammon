package service

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"time"
)

func (s *Service) NewMatch(playerOneID, playerTwoID, length int64) (int64, error) {
	match := new(storage.Match)
	match.PlayerOneID = sql.NullInt64{playerOneID, true}
	match.PlayerTwoID = sql.NullInt64{playerTwoID, true}
	match.Length = sql.NullInt64{length, true}
	match.DateCreated = sql.NullTime{time.Now(), true}

	return s.db.NewMatch(match)

}
