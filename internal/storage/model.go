package storage

import (
	"database/sql"
)

type Player struct {
	ID         sql.NullInt64  `db:"player_id"`
	FirstName  sql.NullString `db:"first_name"`
	LastName   sql.NullString `db:"last_name"`
	TotalGames sql.NullInt64  `db:"total_games"`
	WinGames   sql.NullInt64  `db:"win_games"`
	LoseGames  sql.NullInt64  `db:"loses_games"`
	Experience sql.NullInt64  `db:"experience"'`
	WinsByMars sql.NullInt64  `db:"wins_by_mars"`
	LoseByMars sql.NullInt64  `db:"lose_by_mars"`
	ELORating  sql.NullInt64  `db:"elo_rating"`
	TotalPrize sql.NullInt64  `db:"total_prize"` //TODO
}

type Match struct {
	ID              sql.NullInt64  `db:"id"`
	Length          sql.NullInt64  `db:"length"`
	PlayerOnePoints sql.NullInt64  `db:"player_one_score"`
	PlayerTwoPoints sql.NullInt64  `db:"player_two_score"`
	Status          sql.NullString `db:"status"`
	Rounds          []*Round       `db:"rounds"`
	PlayerOne       *Player        `db:"player_one_id"`
	PlayerTwo       *Player        `db:"player_two_id"`
	DateCreated     sql.NullTime   `db:"date_created"`
	DateUpdated     sql.NullTime   `db:"date_updated"`
}

type Round struct {
	ID       sql.NullInt64 `db:"id"`
	MatchID  sql.NullInt64 `db:"match_id"`
	WinnerID sql.NullInt64 `db:"winner_id"`
	LoserID  sql.NullInt64 `db:"loser_id"`
	IsMars   sql.NullBool  `db:"is_mars"`
	Date     sql.NullTime  `db:"date"`
}

type Tournament struct {
	ID          sql.NullInt64  `db:"id"`
	Players     []*Player      `db:"players"`
	WinnerID    sql.NullInt64  `db:"winner_id"`
	Status      sql.NullString `db:"status"`
	DateCreated sql.NullTime   `db:"date_created"`
	DateUpdated sql.NullTime   `db:"date_updated"`
}
