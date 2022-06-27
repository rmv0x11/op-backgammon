package storage

import (
	"database/sql"
	"github.com/rmv0x11/op-backgammon/internal/model"
	"log"
	"time"
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
		"wins_by_mars" integer,
		"loses_by_mars" integer,
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
	insertPlayer := `insert into players(first_name, last_name) VALUES (?, ?);
		select last_insert_rowid();`

	stmt, err := d.db.Prepare(insertPlayer)
	if err != nil {
		log.Fatalln("insert player prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(player.FirstName, player.LastName)
	if err != nil {
		log.Fatalln("insert player exec error:", err.Error())
		return 0, err
	}

	playerID, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("can't get match_id, err: ", err.Error())
	}
	log.Println("inserting new player record...")

	return playerID, err
}

func (d *Database) GetPlayers() ([]*model.Player, error) {
	rows, err := d.db.Query("select * from players")
	if err != nil {
		log.Fatal("get players query error:", err.Error())
	}
	defer rows.Close()

	players := make([]*model.Player, 0)
	for rows.Next() {
		player := new(model.Player)

		scanErr := rows.Scan(
			&player.ID,
			&player.FirstName,
			&player.LastName,
			&player.TotalGames,
			&player.WinGames,
			&player.LoseGames,
			&player.WinsByMars,
			&player.LoseByMars,
			&player.ELORating,
			&player.TotalPrize,
		)

		if scanErr != nil {
			log.Fatalln("get players scan error:", scanErr)
			return nil, scanErr
		}
		players = append(players, player)
	}

	return players, nil
}

func (d *Database) GetPlayer(id int64) (*model.Player, error) {
	row, err := d.db.Query(`select * from players
		where player_id = ?`,
		id,
	)
	defer row.Close()

	if err != nil {
		log.Fatal("get player info query error:", err.Error())
	}

	player := new(model.Player)

	scanErr := row.Scan(
		&player.ID,
		&player.FirstName,
		&player.LastName,
		&player.TotalGames,
		&player.WinGames,
		&player.LoseGames,
		&player.WinsByMars,
		&player.LoseByMars,
		&player.ELORating,
		&player.TotalPrize,
	)

	if scanErr != nil {
		log.Fatalln("get player info scan error:", scanErr)
		return nil, scanErr
	}

	return player, nil
}

func (d *Database) CreateMatchesTable() error {
	createTable := `create table if not exists matches (
		id integer not null primary key autoincrement,
		length int,
		status text,
		rounds text,
		player_one_id int,
		player_two_id int,
		date timestamp
		);`

	log.Println("create matches table...")

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("create matches table prepare error:", err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("create matches table exec error:", err.Error())
		return err
	}

	log.Println("matches table created")

	return err
}

func (d *Database) CreateRoundsTable() error {
	createTable := `create table if not exists rounds (
		id integer not null primary key autoincrement,
		match_id integer,
		winner_id int,
		"is_mars" boolean not null check (is_mars in (0,1)),
		foreign key(match_id) references matches(id)
		);`

	log.Println("create rounds table...")

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("create rounds table prepare error:", err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("create rounds table exec error:", err.Error())
		return err
	}

	log.Println("rounds table created")

	return err
}

func (d *Database) NewMatch(m *Match) (int64, error) {
	query := `insert into matches(player_one_id, player_two_id, length, date) values (?, ?, ?, ?);
				select last_insert_rowid();`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Fatalln("new match prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(
		m.PlayerOne.ID,
		m.PlayerTwo.ID,
		m.Length,
		m.Date,
	)
	if err != nil {
		log.Fatalln("new match exec error:", err.Error())
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
	query := `insert into rounds(match_id, winner_id,loser_id, is_mars) values (?, ?, ?, ?);
		select last_insert_rowid();`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Fatalln("new round prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(
		r.MatchID,
		r.WinnerID,
		r.LoserID,
		r.IsMars,
	)
	if err != nil {
		log.Fatalln("new round exec error:", err.Error())
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
		date timestamp)`

	log.Println("create tournaments table...")

	stmt, err := d.db.Prepare(createTable)
	if err != nil {
		log.Fatalln("create tournaments table prepare error:", err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln("create tournaments table exec error:", err.Error())
		return err
	}

	log.Println("tournaments table created")

	return err
}

func (d *Database) NewTournament(IDs string) (int64, error) {
	query := `insert into tournaments(players, status, date) values (?, ?, ?);
		select last_insert_rowid();`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Fatalln("new tournament prepare error:", err.Error())
		return 0, err
	}

	res, err := stmt.Exec(
		IDs,
		"Tournament created",
		time.Now(),
	)
	if err != nil {
		log.Fatalln("new round exec error:", err.Error())
		return 0, err
	}

	tournamentID, err := res.LastInsertId()
	if err != nil {
		log.Fatalln("can't get tournament_id, err: ", err.Error())
	}

	log.Println("new tournament added...")

	return tournamentID, err
}

func (d *Database) UpdatePlayer(p *Player) error {
	query := `update players
		set total_games = ?,
			win_games = ?,
			lose_games = ?,
			wins_by_mars = ?,
			lose_by_mars = ?,
			elo_rating = ?,
			total_prize = ?
		where id = ?;`

	stmt, err := d.db.Prepare(query)
	if err != nil {
		log.Fatalln("update player prepare error:", err.Error())
		return err
	}

	_, err = stmt.Exec(
		p.TotalGames,
		p.WinGames,
		p.LoseGames,
		p.WinsByMars,
		p.LoseByMars,
		p.ELORating,
		p.TotalPrize,
	)
	if err != nil {
		log.Fatalln("update player exec error:", err.Error())
		return err
	}

	return nil
}
