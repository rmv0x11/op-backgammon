package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func (a *Application) GetPlayers(c *gin.Context) {
	players, err := a.svc.GetPlayers()
	if err != nil {
		log.Fatalln("DisplayPlayers error:", err)
	}

	for _, v := range players {
		fmt.Printf("player_id:%2d| first_name:%8s| last_name:%s\n",
			v.ID.Int64,
			v.FirstName.String,
			v.LastName.String,
		)
	}
}
