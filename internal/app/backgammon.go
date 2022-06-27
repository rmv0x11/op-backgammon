package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rmv0x11/op-backgammon/internal/service"
	"log"
)

type Application struct {
	svc *service.Service
	r   *gin.Engine
}

func NewBackgammonApp() *Application {
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

	if err := svc.CreateTournamentsTable(); err != nil {
		log.Fatalln("CreateTournamentsTable, err: ", err)
	}

	return &Application{svc: svc, r: gin.New()}
}

func (a *Application) Close() error {
	return a.Close()
}
