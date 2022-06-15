package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/rmv0x11/op-backgammon/internal/app"
	"log"
	"os"
)

func main() {
	//ctx := context.Background()

	os.Remove("backgammon.db")

	impl := app.NewBackgammonAPI()
	defer impl.Close()

	router := gin.New()
	router.POST("/add_player", impl.AddPlayer)
	router.POST("/create_players_table", impl.CreatePlayersTable)
	router.GET("/get_players", impl.GetPlayers)

	err := router.Run(":1337") //TODO port moved into config
	if err != nil {
		log.Fatalf("unable to start the service on port 1337 , error:%s", err.Error())
		return
	}
}
