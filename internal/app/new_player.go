package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"log"
)

func (i *Implementation) NewPlayer(c *gin.Context) {
	firstName := c.Request.URL.Query().Get("firstName")
	lastName := c.Request.URL.Query().Get("lastName")

	playerID, err := i.svc.NewPlayer(&storage.Player{
		FirstName: sql.NullString{firstName, true},
		LastName:  sql.NullString{lastName, true}},
	)
	if err != nil {
		log.Fatalln("AddPlayer handler error:", err.Error())
	}
	log.Println("new match has id:", playerID)
}
