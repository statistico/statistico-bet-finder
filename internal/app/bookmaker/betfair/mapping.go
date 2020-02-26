package betfair

import (
	"fmt"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

func parseCompetitionMapping(id uint64) (string, error) {
	competitions := map[uint64]string{
		8: "10932509",
	}

	if val, ok := competitions[id]; ok {
		return val, nil
	}

	return "", fmt.Errorf("competition ID %d is not supported", id)
}

func parseTeamMapping(team string) string {
	teams := map[string]string{
		"AFC Bournemouth":         "Bournemouth",
		"Brighton & Hove Albion":  "Brighton",
		"Leicester City":          "Leicester",
		"Manchester City":         "Man City",
		"Manchester United":       "Man Utd",
		"Newcastle United":        "Newcastle",
		"Norwich City":            "Norwich",
		"Sheffield United":        "Sheff Utd",
		"Tottenham Hotspur":       "Tottenham",
		"West Ham United":         "West Ham",
		"Wolverhampton Wanderers": "Wolves",
	}

	if val, ok := teams[team]; ok {
		return val
	}

	return team
}

func parseCatalogue(cat []bfClient.MarketCatalogue, fix *proto.Fixture) (*bfClient.MarketCatalogue, error) {
	for _, c := range cat {
		if fixtureMatchesEvent(fix, c.Event) {
			return &c, nil
		}
	}

	return nil, fmt.Errorf("unable to parse event from betfair market catalogues for fixture %d", fix.Id)
}

func fixtureMatchesEvent(fix *proto.Fixture, event bfClient.Event) bool {
	home := parseTeamMapping(fix.HomeTeam.Name)
	away := parseTeamMapping(fix.AwayTeam.Name)

	name := fmt.Sprintf("%s v %s", home, away)

	return name == event.Name
}
