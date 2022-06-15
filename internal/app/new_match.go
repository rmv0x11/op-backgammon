package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (i *Implementation) NewMatch(c *gin.Context) {
	playerOneID := c.Request.URL.Query().Get("player_one_id")
	playerTwoID := c.Request.URL.Query().Get("player_two_id")

	err := i.svc.NewMatch(playerOneID, playerTwoID)
	if err != nil {
		log.Fatalln("unable get info about new match")
	}
}
