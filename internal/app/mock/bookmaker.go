package mock

import (
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/mock"
)

type RunnerFactory struct {
	mock.Mock
}

func (r RunnerFactory) CreateRunner(selectionID uint64, marketID, name string) (*bookmaker.Runner, error) {
	args := r.Called(selectionID, marketID, name)
	b := args.Get(0).(*bookmaker.Runner)
	return b, args.Error(1)
}

type MarketFactory struct {
	mock.Mock
}

func (m MarketFactory) FixtureAndMarket(fix statistico.Fixture, market string) (*bookmaker.Market, error) {
	args := m.Called(fix, market)
	return args.Get(0).(*bookmaker.Market), args.Error(1)
}

type Bookmaker struct {
	mock.Mock
}

func (b Bookmaker) CreateBook(q *app.BookQuery) (*app.Book, error) {
	args := b.Called(q)
	return args.Get(0).(*app.Book), args.Error(1)
}
