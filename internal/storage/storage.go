package storage

import (
	"context"
	"database/sql"
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
	createTable := `CREATE TABLE IF NOT EXISTS players (
		"player_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"first_name" TEXT,
		"last_name" TEXT,
		"total_games" integer,
		"win_games" integer,
		"lose_games" integer,
		"mars_games" integer,
		"elo_rating" integer,
		"total_prize" integer
		);`
	log.Println("Create players table...")

	stmt, err := d.Prepare(createTable)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()

	log.Println("players table created")
}

func (d *Database) InsertPlayer(player *Player) {
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

func (d *Database) GetPlayers() ([]*Player, error) {
	rows, err := d.Query("SELECT * FROM players")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	players := make([]*Player, 0)
	for rows.Next() {
		player := new(Player)
		scanErr := rows.Scan(
			&player.PlayerID,
			&player.FirstName,
			&player.LastName,
			&player.TotalGames,
			&player.WinGames,
			&player.LoseGames,
			&player.MarsGames,
			&player.ELORating,
			&player.TotalPrize,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		players = append(players, player)
	}
	return players, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
