package statistico

import "time"

type Fixture struct {
	ID       uint64    `json:"id"`
	HomeTeam string    `json:"home_team"`
	AwayTeam string    `json:"away_team"`
	Date     time.Time `json:"date"`
}
