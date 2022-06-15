package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/rmv0x11/op-backgammon/internal/app"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	os.Remove("backgammon.db")

	impl := app.NewBackgammonAPI()
	defer impl.Close()

	router := gin.Default()
	router.POST("/add_player", impl.AddPlayer())
	router.Run(":1337")

	err := impl.DisplayPlayers()
	if err != nil {
		log.Fatalln("cannot display players, err:", err)
	}

}

func AddPlayer(c *gin.Context) {

}
