package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"github.com/rmv0x11/op-backgammon/internal/util"
	"log"
	"strconv"
	"time"
)

func (a *Application) UpdateMatch(c *gin.Context) {
	IDValue := c.Request.URL.Query().Get("id")
	playerOnePointsValue := c.Request.URL.Query().Get("player_one_points")
	playerTwoPointsValue := c.Request.URL.Query().Get("player_two_points")
	roundsValue := c.Request.URL.Query().Get("rounds")

	ID, err := strconv.ParseInt(IDValue, 64, 10)
	if err != nil {
		log.Fatalln("can't parse id query params, err: ", err)

	}
	playerOnePoints, err := strconv.ParseInt(playerOnePointsValue, 64, 10)
	if err != nil {
		log.Fatalln("can't parse player_one_points query params, err: ", err)

	}
	playerTwoPoints, err := strconv.ParseInt(playerTwoPointsValue, 64, 10)
	if err != nil {
		log.Fatalln("can't parse player_two_points query params, err :", err)

	}

	m := new(model.Match)
	m.ID = ID
	m.PlayerOnePoints = playerOnePoints
	m.PlayerTwoPoints = playerTwoPoints
	m.DateUpdated = time.Now()
	m.Rounds = util.IDsIntoRounds(roundsValue)

	err = a.svc.UpdateMatch(m)
	if err != nil {
		log.Fatalln("unable update match, err: ", err)
	}
}
