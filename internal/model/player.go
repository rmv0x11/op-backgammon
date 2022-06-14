package model

type Player struct {
	PlayerID   int    `db:"player_id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	TotalGames int    `db:"total_games"`
	WinGames   int    `db:"win_games"`
	LoseGames  int    `db:"lose_games"`
	MarsGames  int    `db:"mars_games"`
	ELORating  int    `db:"elo_rating"`
	TotalPrize int    `db:"total_prize"`
}

type Stats struct {
	Games     int
	Win       int
	Lose      int
	Mars      int
	RatingELO int
}
