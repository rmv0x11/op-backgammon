package app

import "github.com/rmv0x11/op-backgammon/internal/model"

func (i *Implementation) AddPlayer(firstName, lastName string) {
	i.db.InsertPlayer(&model.Player{FirstName: firstName, LastName: lastName})
}
