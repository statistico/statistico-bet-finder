package betfair

import (
	"github.com/statistico/statistico-bet-finder/internal/statistico"
	"github.com/statistico/statistico-bet-finder/internal/statistico/bookmaker"
)

type Service struct {
	BookCreator
}

func (s Service) Book(query bookmaker.ServiceQuery) <-chan *bookmaker.Book {
	ch := make(chan *bookmaker.Book, len(query.Fixtures))

	go s.parseFixtures(query.Fixtures, query.BetTypes, ch)

	return ch
}

func (s Service) parseFixtures(fixtures []statistico.Fixture, types []string, ch chan<- *bookmaker.Book) {
	for _, fix := range fixtures {
		book, err := s.BookCreator.CreateForFixture(fix, types)

		if err != nil {
			// Log error here
		}

		ch <- book
	}

	close(ch)
}