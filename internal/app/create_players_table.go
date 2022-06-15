package app

import (
	"github.com/gin-gonic/gin"
)

func (i *Implementation) CreatePlayersTable(c *gin.Context) {
	i.svc.CreatePlayersTable(c)
}
