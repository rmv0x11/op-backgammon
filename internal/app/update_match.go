package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (a *Application) UpdateMatch(c *gin.Context) {
	err := a.svc.UpdateMatch()
	if err != nil {
		log.Fatalln("unable update match, err: ", err)
	}
}
