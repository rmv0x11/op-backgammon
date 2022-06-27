package service

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/mappers"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"log"
	"math"
)

func (s *Service) NewRound(matchID, winnerID, loserID int64, isMars bool) (int64, error) {
	r := new(storage.Round)
	r.MatchID = sql.NullInt64{matchID, true}
	r.WinnerID = sql.NullInt64{winnerID, true}
	r.LoserID = sql.NullInt64{loserID, true}
	r.IsMars = sql.NullBool{isMars, true}

	{
		winner, err := s.db.GetPlayer(winnerID)
		if err != nil {
			log.Fatalln("can't get winner info, err:", err)
		}

		updWinner := mappers.PlayersForDB(updatePlayer(winner, true, isMars))

		if err = s.db.UpdatePlayer(updWinner); err != nil {
			log.Fatalln("unable update winner, err:", err)
		}
	}

	{
		loser, err := s.db.GetPlayer(loserID)
		if err != nil {
			log.Fatalln("can't get winner info, err:", err)
		}

		updLoser := mappers.PlayersForDB(updatePlayer(loser, false, isMars))

		if err = s.db.UpdatePlayer(updLoser); err != nil {
			log.Fatalln("unable update loser, err:", err)
		}
	}
	return s.db.NewRound(r)
}

func updatePlayer(p *model.Player, isWinner, isMars bool) *model.Player {
	switch isWinner {
	case true:
		if isMars {
			p.WinsByMars++
		}
		p.WinGames++
	case false:
		if isMars {
			p.LoseByMars++
		}
		p.LoseGames++
	}
	p.TotalGames++

	return p
}

//W = (1 – P) * M * S
//
//L = P * M * S , где:
//P = 1 / (1 + pow(10,(-D * sqrt(N) / 2000)))
func updateELO(p, m, s) int64 {

}

func calculateProbability(length, diff float64) float64 {
	return 1 / (1 + math.Pow(10, -diff*math.Sqrt(length)/2000))
}
