package mock

import (
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/mock"
)

type FixtureClient struct {
	mock.Mock
}

func (f FixtureClient) FixtureByID(id uint64) (*statistico.Fixture, error) {
	args := f.Called(id)
	return args.Get(0).(*statistico.Fixture), args.Error(1)
}
