package bootstrap

import (
	"github.com/statistico/statistico-price-finder/internal/app/grpc"
)

func (c Container) GRPCFixtureClient() grpc.FixtureClient {
	return grpc.NewFixtureClient(c.FixtureClient)
}

func (c Container) GRPCOddsCompilerClient() grpc.OddsCompilerClient {
	return grpc.NewOddsCompilerClient(c.OddsCompilerClient)
}
