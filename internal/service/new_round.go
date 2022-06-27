package service

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/mappers"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"log"
)

func (s *Service) NewRound(matchID, winnerID, loserID int64, isMars bool) (int64, error) {
	r := new(storage.Round)
	r.MatchID = sql.NullInt64{matchID, true}
	r.WinnerID = sql.NullInt64{winnerID, true}
	r.LoserID = sql.NullInt64{loserID, true}
	r.IsMars = sql.NullBool{isMars, true}

	//получаем стату победителя
	winner, err := s.db.GetPlayer(winnerID)
	if err != nil {
		log.Fatalln("can't get winner info, err:", err)
	}
	//обновляем его стату
	updWinner := updatePlayerStats(winner, true, isMars)
	//обновляем в бд
	if err = s.db.UpdatePlayer(updWinner); err != nil {
		log.Fatalln("unable update winner, err:", err)
	}
	//получаем стату проигравшего
	loser, err := s.db.GetPlayer(loserID)
	if err != nil {
		log.Fatalln("can't get winner info, err:", err)
	}
	//обновляем его стату
	updLoser := updatePlayerStats(loser, false, isMars)
	//обновляем в бд
	if err = s.db.UpdatePlayer(updLoser); err != nil {
		log.Fatalln("unable update loser, err:", err)
	}

	return s.db.NewRound(r)
}

func updatePlayerStats(p *model.Player, isWinner, isMars bool) *storage.Player {
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

	return mappers.PlayersForDB(p)
}
