package mappers

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

func PlayersForDB(p *model.Player) *storage.Player {
	return &storage.Player{
		ID:         sql.NullInt64{p.ID, true},
		FirstName:  sql.NullString{p.FirstName, true},
		LastName:   sql.NullString{p.LastName, true},
		TotalGames: sql.NullInt64{p.TotalGames, true},
		WinGames:   sql.NullInt64{p.WinGames, true},
		LoseGames:  sql.NullInt64{p.LoseGames, true},
		Experience: sql.NullInt64{p.Experience, true},
		WinsByMars: sql.NullInt64{p.WinsByMars, true},
		LoseByMars: sql.NullInt64{p.LoseByMars, true},
		ELORating:  sql.NullInt64{p.ELORating, true},
		TotalPrize: sql.NullInt64{p.TotalPrize, true},
	}
}

func MatchesForDB(m *model.Match) *storage.Match {
	return &storage.Match{
		ID:              sql.NullInt64{},
		Length:          sql.NullInt64{},
		PlayerOnePoints: sql.NullInt64{},
		PlayerTwoPoints: sql.NullInt64{},
		Status:          sql.NullString{},
		Rounds:          m.Rounds,
		PlayerOne:       nil,
		PlayerTwo:       nil,
		DateCreated:     sql.NullTime{},
		DateUpdated:     sql.NullTime{},
	}
}