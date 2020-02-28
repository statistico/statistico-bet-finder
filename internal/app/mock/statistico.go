package mock

import (
	"github.com/statistico/statistico-price-finder/internal/app"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/statistico/statistico-price-finder/internal/app/statistico"
	"github.com/stretchr/testify/mock"
)

type StatisticoMarketBuilder struct {
	mock.Mock
}

func (m StatisticoMarketBuilder) FixtureAndMarket(f *proto.Fixture, market string) (*statistico.Market, error) {
	args := m.Called(f, market)
	return args.Get(0).(*statistico.Market), args.Error(1)
}

type StatisticoBookmaker struct {
	mock.Mock
}

func (b StatisticoBookmaker) CreateBook(q *app.BookQuery) (*statistico.Book, error) {
	args := b.Called(q)
	return args.Get(0).(*statistico.Book), args.Error(1)
}
