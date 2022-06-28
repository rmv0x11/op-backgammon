package service

import (
	"github.com/rmv0x11/op-backgammon/internal/mappers"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"log"
	"math"
	"time"
)

func (s *Service) NewMatch(playerOneID, playerTwoID, length int64) (int64, error) {
	m := new(model.Match)
	m.PlayerOne.ID = playerOneID
	m.PlayerTwo.ID = playerTwoID
	m.Length = length
	m.DateCreated = time.Now()

	//получить инфу об игроках
	playerOne, err := s.db.GetPlayer(playerOneID)
	if err != nil {
		log.Fatalln("unable get player one, err: ", err)
	}
	playerTwo, err := s.db.GetPlayer(playerTwoID)
	if err != nil {
		log.Fatalln("unable get player two, err: ", err)
	}
	//посчитать изменение эло для первого игрока
	oneWin, oneLose := calculateELO(float64(m.Length), float64(-(playerOne.ELORating - playerTwo.ELORating)), float64(playerOne.Experience))
	m.PlayerOneWin = int64(oneWin)
	m.PlayerOneLose = int64(oneLose)
	//посчитать изменение эло для второго игрока
	twoWin, twoLose := calculateELO(float64(m.Length), float64(-(playerTwo.ELORating - playerOne.ELORating)), float64(playerTwo.Experience))
	m.PlayerTwoWin = int64(twoWin)
	m.PlayerTwoLose = int64(twoLose)

	return s.db.NewMatch(mappers.MatchesForDB(m))

}

func calculateELO(length, diff, exp float64) (win, lose float64) {
	prob := calculateProbability(length, diff)
	var boost float64

	if exp > 400 {
		boost = 1
	} else {
		boost = float64((500 - exp) / 100)
	}
	points := 4 * math.Sqrt(length)

	win = (1 - prob) * boost * points
	lose = prob * boost * points
	return
}

func calculateProbability(length, diff float64) float64 {
	return 1 / (1 + math.Pow(10, -diff*math.Sqrt(length)/2000))
}
