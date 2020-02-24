package app

import (
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
)

type BookMaker interface {
	CreateBook(q *BookQuery) (*Book, error)
}

type BookQuery struct {
	EventID  	uint64
	Markets    []string
}

// BookMaker is responsible for creating a Book struct of bookmaker markets.
type bookMaker struct {
	fixtureClient statistico.FixtureClient
	builder       MarketBuilder
	clock         clockwork.Clock
	logger        *logrus.Logger
}

// CreateBook creates a Book struct of Statistico and Bookmaker markets.
func (b bookMaker) CreateBook(q *BookQuery) (*Book, error) {
	book := Book{
		EventID:   q.EventID,
		CreatedAt: b.clock.Now(),
	}

	fixture, err := b.fixtureClient.FixtureByID(q.EventID)

	if err != nil {
		return &book, errNotFound
	}

	for _, m := range q.Markets {
		market, err := b.builder.FixtureAndMarket(fixture, m)

		if err != nil {
			b.logger.Warn(err.Error())
			continue
		}

		book.Markets = append(book.Markets, market)
	}

	return &book, nil
}

func NewBookMaker(f statistico.FixtureClient, b MarketBuilder, c clockwork.Clock, l *logrus.Logger) BookMaker {
	return &bookMaker{
		fixtureClient: f,
		builder:       b,
		clock:         c,
		logger:        l,
	}
}
