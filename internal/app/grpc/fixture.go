package grpc

import (
	"context"
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
	"time"
)

type FixtureClient struct {
	client proto.FixtureServiceClient
}

func (d FixtureClient) FixtureByID(id uint64) (*app.Fixture, error) {
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
