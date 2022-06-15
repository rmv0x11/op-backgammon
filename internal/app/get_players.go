package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func (i *Implementation) GetPlayers(c *gin.Context) {
	players, err := i.svc.GetPlayers(c)
	if err != nil {
		log.Fatalln("DisplayPlayers error:", err)
	}

	for _, v := range players {
		fmt.Printf("player_id:%2d| first_name:%8s| last_name:%s\n",
			v.PlayerID.Int64,
			v.FirstName.String,
			v.LastName.String,
		)
	}
}
