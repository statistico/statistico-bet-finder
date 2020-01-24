package statistico

import "time"

// Fixture is a struct containing fixture specific information
type Fixture struct {
	ID            uint64    `json:"id"`
	CompetitionID uint64    `json:"competition_id"`
	HomeTeam      string    `json:"home_team"`
	AwayTeam      string    `json:"away_team"`
	Date          time.Time `json:"date"`
}
