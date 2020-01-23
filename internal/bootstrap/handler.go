package bootstrap

import "github.com/statistico/statistico-bet-finder/internal/app/rest"

func (c Container) BookHandler() *rest.BookHandler {
	return rest.NewBookHandler(c.BookMaker())
}
