package storage

import (
	"database/sql"
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

func (d *Database) NewPlayer(player *Player) (int64, error) {
	insertPlayer := `INSERT INTO players(first_name, last_name) VALUES (?, ?);
		SELECT last_insert_rowid();`

	stmt, err := d.db.Prepare(insertPlayer)
	if err != nil {
		log.Fatalln("InsertPlayer prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(player.FirstName, player.LastName)
	if err != nil {
		log.Fatalln("InsertPlayer exec error:", err.Error())
		return 0, err
	}

	playerID, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("can't get match_id, err: ", err.Error())
	}
	log.Println("inserting new player record...")

	return playerID, err
}

func (d *Database) GetPlayers() ([]*Player, error) {
	rows, err := d.db.Query("SELECT * FROM players")
	if err != nil {
		log.Fatal("GetPlayers query error:", err.Error())
	}
	defer rows.Close()

	players := make([]*Player, 0)
	for rows.Next() {
		player := new(Player)

		scanErr := rows.Scan(
			&player.ID,
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
		&player.ID,
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

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("CreateRoundsTable exec error:", err.Error())
		return err
	}

	log.Println("rounds table created")

	return err
}

func (d *Database) NewMatch(m *Match) (int64, error) {
	query := `INSERT INTO matches(player_one_id, player_two_id, length, date) VALUES (?, ?, ?, ?);
				SELECT last_insert_rowid();`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Fatalln("NewMatch prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(
		m.PlayerOne.ID,
		m.PlayerTwo.ID,
		m.Length,
		m.Date,
	)
	if err != nil {
		log.Fatalln("NewMatch exec error:", err.Error())
		return 0, err
	}

	matchID, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("can't get match_id, err: ", err.Error())
	}

	log.Println("new match created...")

	return matchID, err
}

func (d *Database) NewRound(r *Round) (int64, error) {
	query := `insert into rounds(match_id, winner_id, is_mars) values (?, ?, ?);
		SELECT last_insert_rowid();`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Fatalln("NewRound prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(
		r.MatchID,
		r.WinnerID,
		r.IsMars,
	)
	if err != nil {
		log.Fatalln("NewRound exec error:", err.Error())
		return 0, err
	}

	roundID, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("can't get round_id, err: ", err.Error())
	}

	log.Println("new round added...")

	return roundID, err

}

func (d *Database) CreateTournamentTables() error {
	createTable := `create table if not exists tournaments (
		id integer not null primary key autoincrement,
		players text,
		winner_id integer,
		status text,
		date timestamp;
		select last_insert_rowid();'`

	log.Println("Create tournaments table...")

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("CreateTournamentsTable prepare error:", err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("CreateTournamentsTable exec error:", err.Error())
		return err
	}

	log.Println("tournaments table created")

	return err
}
