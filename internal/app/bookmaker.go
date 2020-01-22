package app

import (
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type BookMaker interface {
	CreateBook(q *BookQuery) *Book
}

type BookQuery struct {
	Markets    []string
	FixtureIDs []uint64
}

// BookMaker is responsible for creating a Book struct of Statistico and Bookmaker markets.
type bookMaker struct {
	fixtureClient statistico.FixtureClient
	builder       MarketBuilder
	clock         clockwork.Clock
	logger        *logrus.Logger
}

// CreateBook creates a Book struct of Statistico and Bookmaker markets.
func (b bookMaker) CreateBook(q *BookQuery) *Book {
	book := Book{
		Markets:   []*Market{},
		CreatedAt: b.clock.Now(),
	}

	for _, id := range q.FixtureIDs {
		fixture, err := b.fixtureClient.FixtureByID(id)

		if err != nil {
			b.logger.Warnf("Error '%s' fetching fixture '%d' when creating a book", err.Error(), id)
			continue
		}

		for _, m := range q.Markets {
			market, err := b.builder.FixtureAndMarket(fixture, m)

			if err != nil {
				b.logger.Warn(err.Error())
				continue
			}

			book.Markets = append(book.Markets, market)
		}
	}

	return &book
}

func NewBookMaker(f statistico.FixtureClient, b MarketBuilder, c clockwork.Clock, l *logrus.Logger) BookMaker {
	return &bookMaker{
		fixtureClient: f,
		builder:       b,
		clock:         c,
		logger:        l,
	}
}
