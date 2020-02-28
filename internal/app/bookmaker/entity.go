package bookmaker

import "time"

type Book struct {
	EventID   uint64    `json:"eventId"`
	Markets   []*Market `json:"markets"`
	CreatedAt time.Time `json:"createdAt"`
}

type Market struct {
	Name       string       `json:"name"`
	Bookmakers []*SubMarket `json:"bookmakers"`
}

type SubMarket struct {
	ID        string   `json:"marketId"`
	Bookmaker string   `json:"bookmaker"`
	Runners   []Runner `json:"runners"`
}

type Runner struct {
	Name        string  `json:"name"`
	SelectionID uint64  `json:"selectionId"`
	Back        []Price `json:"back"`
	Lay         []Price `json:"lay"`
}

type Price struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}
