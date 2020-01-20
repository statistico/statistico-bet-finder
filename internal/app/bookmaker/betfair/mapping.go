package betfair

import (
	"fmt"
	"github.com/statistico/statistico-bet-finder/internal/app"
	bfClient "github.com/statistico/statistico-betfair-go-client"
)

func parseCompetitionMapping(id uint64) (string, error) {
	competitions := map[uint64]string{
		16036: "10932509",
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

func fixtureMatchesEvent(fix app.Fixture, event bfClient.Event) bool {
	home := parseTeamMapping(fix.HomeTeam)
	away := parseTeamMapping(fix.AwayTeam)

	name := fmt.Sprintf("%s v %s", home, away)

	return name == event.Name
}
