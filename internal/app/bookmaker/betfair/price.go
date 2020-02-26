package betfair

import (
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/bookmaker"
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
