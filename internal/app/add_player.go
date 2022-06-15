package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

func (i *Implementation) AddPlayer(c *gin.Context) error {
	//i.r.POST("/add_player", func(c *gin.Context) {
	//	firstName := c.Request.URL.Query().Get("firstName")
	//	lastName := c.Request.URL.Query().Get("lastName")
	//	err := i.svc.AddPlayer(&storage.Player{FirstName: sql.NullString{firstName}})
	//	if err != nil {
	//		log.Fatalln()
	//		return
	//	}
	//})
	err := i.svc.AddPlayer(&storage.Player{
		FirstName: sql.NullString{firstName, true},
		LastName:  sql.NullString{lastName, true}},
	)

}
