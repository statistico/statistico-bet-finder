package mock

import (
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
