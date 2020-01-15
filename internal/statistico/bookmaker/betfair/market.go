package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

type MarketBuilder struct {
	RunnerCreator
}

func (m MarketBuilder) Build(c *bfClient.MarketCatalogue) (*bookmaker.Market, error) {
	market := bookmaker.Market{
		ID:   c.MarketID,
		Name: c.MarketName,
	}

	for _, run := range c.Runners {
		r, err := m.RunnerCreator.Create(c.MarketID, run)

		if err != nil {
			return nil, err
		}

		market.Runners = append(market.Runners, *r)
	}

	return &market, nil
}
