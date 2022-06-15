package app

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/storage"
)

func (i *Implementation) AddPlayer(firstName, lastName string) {
	i.db.InsertPlayer(&storage.Player{
		FirstName: sql.NullString{firstName, true},
		LastName:  sql.NullString{lastName, true}},
	)
}
