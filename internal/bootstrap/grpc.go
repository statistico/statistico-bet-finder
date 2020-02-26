package bootstrap

import (
	"github.com/statistico/statistico-price-finder/internal/app/grpc"
)

func (c Container) GRPCFixtureClient() grpc.FixtureClient {
	return grpc.NewFixtureClient(c.FixtureClient)
}
