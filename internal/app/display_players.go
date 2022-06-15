package app

import (
	"fmt"
	"log"
)

func (i *Implementation) DisplayPlayers() error {
	players, err := i.db.GetPlayers()
	if err != nil {
		log.Fatalln("DisplayPlayers error:", err.Error())
		return err
	}

	for _, v := range players {
		fmt.Printf("player_id:%2d| first_name:%8s| last_name:%s\n",
			v.PlayerID.Int64,
			v.FirstName.String,
			v.LastName.String,
		)
	}
	return err
}
