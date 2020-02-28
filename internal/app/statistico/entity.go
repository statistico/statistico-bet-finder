package statistico

import "time"

type Book struct {
	EventID   uint64    `json:"eventId"`
	Markets   []*Market `json:"markets"`
	CreatedAt time.Time `json:"createdAt"`
}

type Market struct {
	Name    string    `json:"name"`
	Runners []*Runner `json:"runners"`
}

type Runner struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
