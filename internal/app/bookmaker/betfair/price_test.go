package betfair

import (
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildPrices(t *testing.T) {
	t.Run("returns hydrated slice of price", func(t *testing.T) {
		betfair := []bfClient.PriceSize{
			{
				Price: 1.89,
				Size:  56.90,
			},
			{
				Price: 10.94,
				Size:  560.90,
			},
			{
				Price: 15.81,
				Size:  6.90,
			},
		}

		prices := buildPrices(betfair)

		assert.Equal(t, float32(1.89), prices[0].Price)
		assert.Equal(t, float32(56.90), prices[0].Size)
		assert.Equal(t, float32(10.94), prices[1].Price)
		assert.Equal(t, float32(560.90), prices[1].Size)
		assert.Equal(t, float32(15.81), prices[2].Price)
		assert.Equal(t, float32(6.90), prices[2].Size)
	})
}