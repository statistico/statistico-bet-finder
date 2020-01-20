package app

import (
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type BookQuery struct {
	BetTypes []string
	FixtureIDs []uint64
}

type BookFactory struct {
	fixtureClient statistico.FixtureClient
	builder MarketBuilder
	// Add clock implementation here
}

func (b BookFactory) CreateBook(q BookQuery) *Book {
	// Create book with time set here
	book := Book{}

	for _, id := range q.FixtureIDs {
		fixture, err := b.fixtureClient.FixtureByID(id)

		if err != nil {
			// Log error here
			continue
		}

		for _, t := range q.BetTypes {
			market := b.builder.FixtureAndBetType(fixture, t)

			if market == nil {
				continue
			}

			book.Market = append(book.Market, market)
		}
	}

	return &book
}
