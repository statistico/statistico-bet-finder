package mock

import (
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
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
