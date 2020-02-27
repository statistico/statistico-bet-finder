package mock

import (
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
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

func (m MarketFactory) FixtureAndMarket(fix *proto.Fixture, market string) (*bookmaker.SubMarket, error) {
	args := m.Called(fix, market)
	return args.Get(0).(*bookmaker.SubMarket), args.Error(1)
}

type Bookmaker struct {
	mock.Mock
}

func (b Bookmaker) CreateBook(q *bookmaker.BookQuery) (*bookmaker.Book, error) {
	args := b.Called(q)
	return args.Get(0).(*bookmaker.Book), args.Error(1)
}

type MarketBuilder struct {
	mock.Mock
}

func (m MarketBuilder) FixtureAndMarket(f *proto.Fixture, market string) (*bookmaker.Market, error) {
	args := m.Called(f, market)
	return args.Get(0).(*bookmaker.Market), args.Error(1)
}
