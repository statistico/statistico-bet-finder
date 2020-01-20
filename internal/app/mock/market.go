package mock

import (
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/stretchr/testify/mock"
)

type MarketBuilder struct {
	mock.Mock
}

func (m MarketBuilder) FixtureAndBetType(f *app.Fixture, bet string) *app.Market {
	args := m.Called(f, bet)
	return args.Get(0).(*app.Market)
}
