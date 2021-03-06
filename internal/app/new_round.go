package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (a *Application) NewRound(c *gin.Context) {
	matchIDValue := c.Request.URL.Query().Get("match_id")
	winnerIDValue := c.Request.URL.Query().Get("winner_id")
	loserIDValue := c.Request.URL.Query().Get("loser_id")

	isMarsValue := c.Request.URL.Query().Get("is_mars")

	matchID, err := strconv.ParseInt(matchIDValue, 10, 64)
	if err != nil {
		log.Fatalln("can't parse match_id query params, err: ", err)
	}
	winnerID, err := strconv.ParseInt(winnerIDValue, 10, 64)
	if err != nil {
		log.Fatalln("can't parse winner_id query params, err: ", err)
	}
	loserID, err := strconv.ParseInt(loserIDValue, 10, 64)
	if err != nil {
		log.Fatalln("can't parse winner_id query params, err: ", err)
	}
	isMars, err := strconv.ParseBool(isMarsValue)
	if err != nil {
		log.Fatalln("can't parse is_mars query params, err: ", err)
	}

	roundID, err := a.svc.NewRound(matchID, winnerID, loserID, isMars)
	if err != nil {
		log.Fatalln("unable added new round, err: ", err)
	}

	_, err = c.Writer.Write([]byte(strconv.FormatInt(roundID, 10)))
	if err != nil {
		log.Fatalln("can't write response, err: ", err)
	}
}
