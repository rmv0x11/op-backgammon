package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/rmv0x11/op-backgammon/internal/app"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	os.Remove("backgammon.db")

	impl := app.NewBackgammonAPI(ctx)
	defer impl.Close()

	impl.CreatePlayersTable(ctx)

	impl.AddPlayer("john", "cena")
	impl.AddPlayer("mark", "aurelius")
	impl.AddPlayer("viktor", "tsoi")
	// INSERT RECORDS
	err := impl.DisplayPlayers()
	if err != nil {
		log.Fatalln("cannot display players, err:", err)
	}

	// DISPLAY INSERTED RECORDS
	//displayStudents(sqliteDatabase)
}
