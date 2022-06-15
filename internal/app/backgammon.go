package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/service"
)

type Implementation struct {
	svc *service.Service
	r   *gin.Engine
}

func NewBackgammonAPI() *Implementation {
	svc := service.NewService()

	return &Implementation{svc: svc, r: gin.New()}
}

func (i *Implementation) Close() error {
	return i.Close()
}
