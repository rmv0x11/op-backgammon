package model

import (
	"time"
)

type Player struct {
	ID         int64
	FirstName  string
	LastName   string
	TotalGames int64
	WinGames   int64
	LoseGames  int64
	Experience int64
	WinsByMars int64
	LoseByMars int64
	ELORating  int64
	TotalPrize int64 //TODO
}

type Match struct {
	ID              int64
	Length          int64
	PlayerOnePoints int64
	PlayerTwoPoints int64
	Status          string
	Rounds          []*Round
	PlayerOne       *Player
	PlayerTwo       *Player
	DateCreated     time.Time
	DateUpdated     time.Time
}

type Round struct {
	ID       int64
	MatchID  int64
	WinnerID int64
	LoserID  int64
	IsMars   bool
	Date     time.Time
}

type Tournament struct {
	ID          int64
	Players     []*Player
	WinnerID    int64
	Status      string
	DateCreated time.Time
	DateUpdated time.Time
}
