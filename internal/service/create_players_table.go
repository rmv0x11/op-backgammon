package service

import "github.com/gin-gonic/gin"

func (s *Service) CreatePlayersTable(c *gin.Context) {
	s.db.CreatePlayersTable()
}
