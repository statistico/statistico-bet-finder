package grpc

import (
	"context"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

// FixtureClient is a wrapper around the Statistico Data service.
type FixtureClient interface {
	FixtureByID(id uint64) (*proto.Fixture, error)
}

type fixtureClient struct {
	client proto.FixtureServiceClient
}

// FixtureByID returns a fixture struct parsed from the Statistico data service.
func (d fixtureClient) FixtureByID(id uint64) (*proto.Fixture, error) {
	request := proto.FixtureRequest{FixtureId: id}

	response, err := d.client.FixtureByID(context.Background(), &request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func NewFixtureClient(c proto.FixtureServiceClient) FixtureClient {
	return &fixtureClient{client: c}
}
