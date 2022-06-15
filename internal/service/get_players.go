package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

func (s *Service) GetPlayers(c *gin.Context) ([]*storage.Player, error) {
	return s.db.GetPlayers(c)
}
