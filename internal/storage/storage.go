package storage

import (
	"context"
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"log"
)

type Database struct {
	*sql.DB
}

//NewStorage return new Storage
func NewStorage(db *sql.DB) *Database {
	return &Database{db}
}

//CreatePlayersTable creates new table Players in Storage
func (d *Database) CreatePlayersTable(ctx context.Context) {
	createTable := `CREATE TABLE players (
		"player_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"first_name" TEXT,
		"last_name" TEXT,
		"total_games" integer,
		"win_games" integer,
		"lose_games" integer,
		"mars_games" integer,
		"elo_rating" integer,
		"total_prize"
		);`
	log.Println("Create players table...")

	stmt, err := d.Prepare(createTable)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()

	log.Println("players table created")
}

func (d *Database) InsertPlayer(player *model.Player) {
	log.Println("inserting new player record...", player.FirstName)

	insertPlayer := `INSERT INTO players(first_name, last_name) VALUES (?, ?)`

	stmt, err := d.Prepare(insertPlayer)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec(player.FirstName, player.LastName)
	if err != nil {
		log.Fatalln(err)
	}
}

func (d *Database) GetPlayers() []*model.Player {
	row, err := d.Query("SELECT * FROM players")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	players := make([]*model.Player, 0)
	for row.Next() {
		player := new(model.Player)
		row.Scan(&player)
		players = append(players, player)
	}
	return players
}

func (d *Database) Close() error {
	return d.DB.Close()
}
