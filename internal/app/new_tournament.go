package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (a *Application) NewTournament(c *gin.Context) {
	playersIDs := c.Request.URL.Query().Get("players_ids")

	tournamentID, err := a.svc.NewTournament(playersIDs)
	if err != nil {
		log.Fatalln("unable added new tournament, err:", err)
	}

	log.Println("new tournament has id:", tournamentID)
}
