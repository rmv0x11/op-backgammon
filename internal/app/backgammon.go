package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/service"
	"log"
)

type Implementation struct {
	svc *service.Service
	r   *gin.Engine
}

func NewBackgammonAPI() *Implementation {
	svc := service.NewService()

	if err := svc.CreatePlayersTable(); err != nil {
		log.Fatalln("CreatePlayersTable, err: ", err)
	}

	if err := svc.CreateMatchesTable(); err != nil {
		log.Fatalln("CreateMatchesTable, err: ", err)
	}

	if err := svc.CreateRoundsTable(); err != nil {
		log.Fatalln("CreateRoundsTable, err: ", err)
	}

	return &Implementation{svc: svc, r: gin.New()}
}

func (i *Implementation) Close() error {
	return i.Close()
}
