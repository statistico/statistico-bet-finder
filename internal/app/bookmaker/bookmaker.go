package bookmaker

import (
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-price-finder/internal/app"
	"github.com/statistico/statistico-price-finder/internal/app/grpc"
)

type BookMaker interface {
	CreateBook(q *app.BookQuery) (*Book, error)
}

// BookMaker is responsible for creating a Book struct of bookmaker markets.
type bookMaker struct {
	fixtureClient grpc.FixtureClient
	builder       MarketBuilder
	clock         clockwork.Clock
	logger        *logrus.Logger
}

// CreateBook creates a Book struct of Statistico and Bookmaker markets.
func (b bookMaker) CreateBook(q *app.BookQuery) (*Book, error) {
	book := Book{
		EventID:   q.EventID,
		CreatedAt: b.clock.Now(),
	}

	fixture, err := b.fixtureClient.FixtureByID(q.EventID)

	if err != nil {
		return &book, app.ErrNotFound
	}

	for _, m := range q.Markets {
		market, err := b.builder.FixtureAndMarket(fixture, m)

		if err != nil {
			b.logger.Warnf("Error building market for event %d: %s", q.EventID, err.Error())
			continue
		}

		book.Markets = append(book.Markets, market)
	}

	return &book, nil
}

func NewBookMaker(f grpc.FixtureClient, b MarketBuilder, c clockwork.Clock, l *logrus.Logger) BookMaker {
	return &bookMaker{
		fixtureClient: f,
		builder:       b,
		clock:         c,
		logger:        l,
	}
}
