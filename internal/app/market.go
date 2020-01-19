package app

// Market is a struct containing Statistico calculated odds for a fixture
type Market struct {
	FixtureID uint64   `json:"fixture_id"`
	Name      string   `json:"name"`
	Runners []Runner   `json:"runners"`
}

// Runner is a struct containing individual runner information
type Runner struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
}
