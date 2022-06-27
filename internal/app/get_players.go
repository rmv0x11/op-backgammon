package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/util"
	"log"
)

func (a *Application) GetPlayers(c *gin.Context) {
	players, err := a.svc.GetPlayers()
	if err != nil {
		log.Fatalln("DisplayPlayers error:", err)
	}

	playersIDs := util.PlayersIntoIDs(players)

	_, err = c.Writer.Write([]byte(playersIDs))
	if err != nil {
		log.Fatalln("can't write response, err:", err)
	}
}
