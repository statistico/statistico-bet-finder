package statistico_test

import (
	"errors"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/statistico/statistico-price-finder/internal/app/mock"
	"github.com/statistico/statistico-price-finder/internal/app/statistico"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarketBuilder_FixtureAndMarket(t *testing.T) {
	t.Run("returns a hydrated market struct", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.OddsCompilerClient)
		builder := statistico.NewMarketBuilder(mockClient)

		fixture := proto.Fixture{Id: 45381}

		response := proto.EventMarket{
			EventId:              45381,
			Market:               "OVER_UNDER_25",
			Odds:                 []*proto.Odds{
				{
					Selection: "Over 2.5 Goals",
					Price: 1.95,
				},
				{
					Selection: "Under 2.5 Goals",
					Price: 2.06,
				},
			},
		}

		mockClient.On("EventMarket", uint64(45381), "OVER_UNDER_25").Return(&response, nil)

		market, err := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Error building market expected nil got %s", err)
		}

		mockClient.AssertExpectations(t)
		assert.Equal(t, "OVER_UNDER_25", market.Name)
		assert.Equal(t, 2, len(market.Runners))
		assert.Equal(t, "Over 2.5 Goals", market.Runners[0].Name)
		assert.Equal(t, float32(1.95), market.Runners[0].Price)
		assert.Equal(t, "Under 2.5 Goals", market.Runners[1].Name)
		assert.Equal(t, float32(2.06), market.Runners[1].Price)
	})

	t.Run("error is returned if error returned by odds client", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.OddsCompilerClient)
		builder := statistico.NewMarketBuilder(mockClient)

		fixture := proto.Fixture{Id: 45381}

		mockClient.On("EventMarket", uint64(45381), "OVER_UNDER_25").Return(&proto.EventMarket{}, errors.New("error occurred"))

		_, err := builder.FixtureAndMarket(&fixture, "OVER_UNDER_25")

		if err == nil {
			t.Fatal("Error building market expected error got nil")
		}

		assert.Equal(t, "error occurred", err.Error())
	})
}
