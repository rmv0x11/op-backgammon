package app

import (
	"github.com/gin-gonic/gin"

	"log"
	"strconv"
)

func (a *Application) NewMatch(c *gin.Context) {
	playerOneValue := c.Request.URL.Query().Get("player_one_id")
	playerTwoValue := c.Request.URL.Query().Get("player_two_id")
	lengthValue := c.Request.URL.Query().Get("length")

	playerOneID, err := strconv.ParseInt(playerOneValue, 10, 64)
	if err != nil {
		log.Fatalln("can't parse player one query params, err:")
	}
	playerTwoID, err := strconv.ParseInt(playerTwoValue, 10, 64)
	if err != nil {
		log.Fatalln("can't parse player two query params, err: ", err)

	}

	length, err := strconv.ParseInt(lengthValue, 10, 64)
	if err != nil {
		log.Fatalln("can't parse length query params, err:")
	}

	matchID, err := a.svc.NewMatch(playerOneID, playerTwoID, length)
	if err != nil {
		log.Fatalln("unable get info about new match, err: ", err)
	}

	log.Println("new match has id:", matchID)
}
