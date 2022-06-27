package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func (a *Application) NewTournament(c *gin.Context) {
	playersIDs := c.Request.URL.Query().Get("players_ids")

	tournamentID, err := a.svc.NewTournament(playersIDs)
	if err != nil {
		log.Fatalln("unable added new tournament, err:", err)
	}

	_, err = c.Writer.Write([]byte(strconv.FormatInt(tournamentID, 10)))
	if err != nil {
		log.Fatalln("can't write response, err:", err)
	}
}
