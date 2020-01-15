package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
)

type Service struct {
	BookCreator
}

func (s Service) Markets(query bookmaker.ServiceQuery) <-chan *bookmaker.Market {
	ch := make(chan *bookmaker.Market, len(query.Fixtures))

	go s.parseFixtures(query, ch)

	return ch
}

func (s Service) parseFixtures(query bookmaker.ServiceQuery, ch chan<- *bookmaker.Market) {
	for _, fix := range query.Fixtures {
		for _, betType := range query.BetTypes {
			book, err := s.BookCreator.FixtureAndBetType(fix, betType)

			if err != nil {
				// Log error here
			}

			ch <- book
		}
	}

	close(ch)
}
