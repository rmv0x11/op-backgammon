package storage

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

type Database struct {
	db *sql.DB
}

//NewStorage return new Storage
func NewStorage(dsn string) *Database {
	db := getSQLite(dsn)
	return &Database{db}
}

func getSQLite(dsn string) *sql.DB {
	sqliteDB, _ := sql.Open("sqlite3", dsn)
	return sqliteDB
}

func (d *Database) Close() error {
	return d.db.Close()
}

//CreatePlayersTable creates new table Players in Storage
func (d *Database) CreatePlayersTable() error {
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

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("CreatePlayersTable prepare error:", err.Error())
		return err
	}

	log.Println("Create players table...")

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("CreatePlayersTable exec error:", err.Error())
		return err
	}

	log.Println("players table created")

	return err
}

func (d *Database) InsertPlayer(player *Player) error {
	insertPlayer := `INSERT INTO players(first_name, last_name) VALUES (?, ?)`

	stmt, err := d.db.Prepare(insertPlayer)
	if err != nil {
		log.Fatalln("InsertPlayer prepare error:", err.Error())
		return err
	}

	_, err = stmt.Exec(player.FirstName, player.LastName)
	if err != nil {
		log.Fatalln("InsertPlayer exec error:", err.Error())
		return err
	}

	log.Println("inserting new player record...")

	return err
}

func (d *Database) GetPlayers(c *gin.Context) ([]*Player, error) {
	rows, err := d.db.Query("SELECT * FROM players")
	if err != nil {
		log.Fatal("GetPlayers query error:", err.Error())
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
			log.Fatalln("GetPlayers scan error:", scanErr)
			return nil, scanErr
		}
		players = append(players, player)
	}

	return players, nil
}

func (d *Database) GetPlayerInfo(id int) (*Player, error) {
	row, err := d.db.Query(`SELECT * FROM players
		WHERE player_id=?`,
		id,
	)
	defer row.Close()

	if err != nil {
		log.Fatal("GetPlayerInfo query error:", err.Error())
	}

	player := new(Player)

	scanErr := row.Scan(
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
		log.Fatalln("GetPlayerInfo scan error:", scanErr)
		return nil, scanErr
	}

	return player, nil
}

func (d *Database) CreateMatchesTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS matches (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"length" int,
		"status" text,
		"rounds" text,
		"player_one_id" int,
		"player_two_id" int,
		"date" timestamp
		);`

	log.Println("Create matches table...")

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("CreateMatchesTable prepare error:", err.Error())
		return err
	}

	log.Println("Create matches table...")

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("CreateMatchesTable exec error:", err.Error())
		return err
	}

	log.Println("matches table created")

	return err
}

func (d *Database) CreateRoundsTable() error {
	createTable := `CREATE TABLE IF NOT EXISTS rounds (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"match_id" integer,
		"winner_id" int,
		"is_mars" boolean not null check (is_mars in (0,1)),
		foreign key(match_id) references matches(id)
		);`

	log.Println("Create rounds table...")

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("CreateRoundsTable prepare error:", err.Error())
		return err
	}

	log.Println("Create rounds table...")

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("CreateRoundsTable exec error:", err.Error())
		return err
	}

	log.Println("rounds table created")

	return err
}

func (d *Database) NewMatch() {

}
