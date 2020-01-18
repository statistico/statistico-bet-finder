package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/app/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

func buildPrices(prices []bfClient.PriceSize) []bookmaker.Price {
	var p []bookmaker.Price

	for _, price := range prices {
		x := bookmaker.Price{
			Price: price.Price,
			Size:  price.Size,
		}

		p = append(p, x)
	}

	return p
}
