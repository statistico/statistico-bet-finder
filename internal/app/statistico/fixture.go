package statistico

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
	"time"
)

// GRPCFixtureClient is a wrapper around the Statistico Data service.
type GRPCFixtureClient struct {
	client proto.FixtureServiceClient
}

// FixtureByID returns a fixture struct parsed from the Statistico data service.
func (d GRPCFixtureClient) FixtureByID(id uint64) (*app.Fixture, error) {
	request := proto.FixtureRequest{FixtureId: id}

	response, err := d.client.FixtureByID(context.Background(), &request)

	if err != nil {
		return nil, err
	}

	return convertResponseToFixture(response), err
}

func convertResponseToFixture(resp *proto.Fixture) *app.Fixture {
	fixture := app.Fixture{
		ID:            uint64(resp.Id),
		CompetitionID: uint64(resp.Competition.Id),
		HomeTeam:      resp.HomeTeam.Name,
		AwayTeam:      resp.AwayTeam.Name,
		Date:          time.Unix(resp.DateTime.Utc, 0),
	}

	return &fixture
}
