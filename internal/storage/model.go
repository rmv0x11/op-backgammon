package storage

import "database/sql"

type Player struct {
	PlayerID   sql.NullInt64  `db:"player_id"`
	FirstName  sql.NullString `db:"first_name"`
	LastName   sql.NullString `db:"last_name"`
	TotalGames sql.NullInt64  `db:"total_games"`
	WinGames   sql.NullInt64  `db:"win_games"`
	LoseGames  sql.NullInt64  `db:"lose_games"`
	MarsGames  sql.NullInt64  `db:"mars_games"`
	ELORating  sql.NullInt64  `db:"elo_rating"`
	TotalPrize sql.NullInt64  `db:"total_prize"`
}

type Match struct {
	Player1 Player
	Player2 Player
}
