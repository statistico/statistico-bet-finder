package statistico

import "time"

// Market is a struct containing Statistico calculated odds for a fixture.
type Market struct {
	FixtureID uint64   `json:"fixture_id"`
	Name      string   `json:"name"`
	Runners   []Runner `json:"runners"`
}

// Runner is a struct containing individual runner information.
type Runner struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// Fixture is a struct containing fixture specific information
type Fixture struct {
	ID            uint64    `json:"id"`
	CompetitionID uint64    `json:"competition_id"`
	HomeTeam      string    `json:"home_team"`
	AwayTeam      string    `json:"away_team"`
	Date          time.Time `json:"date"`
}
