package storage

import (
	"database/sql"
)

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
	ID        sql.NullInt64  `db:"id"`
	Length    sql.NullInt64  `db:"length"`
	Status    sql.NullString `db:"status"`
	Rounds    []Round        `db:"rounds"`
	PlayerOne Player         `db:"player_one_id"`
	PlayerTwo Player         `db:"player_two_id"`
	Date      sql.NullTime   `db:"date"`
}

type Round struct {
	ID       sql.NullInt64 `db:"id"`
	MatchID  sql.NullInt64 `db:"match_id"`
	WinnerID sql.NullInt64 `db:"winner_id"`
	IsMars   sql.NullBool  `db:"is_mars"`
}
