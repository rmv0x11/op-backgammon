package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"log"
	"strconv"
)

const baseELORating = 1200

func (a *Application) NewPlayer(c *gin.Context) {
	firstName := c.Request.URL.Query().Get("firstName")
	lastName := c.Request.URL.Query().Get("lastName")

	playerID, err := a.svc.NewPlayer(&storage.Player{
		FirstName: sql.NullString{firstName, true},
		LastName:  sql.NullString{lastName, true},
		ELORating: sql.NullInt64{baseELORating, true},
	},
	)
	if err != nil {
		log.Fatalln("AddPlayer handler error:", err.Error())
	}

	_, err = c.Writer.Write([]byte(strconv.FormatInt(playerID, 10)))
	if err != nil {
		log.Fatalln("can't write response, err:", err)
	}
}
