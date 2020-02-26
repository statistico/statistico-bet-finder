package mock

import (
	"context"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type FixtureServiceClient struct {
	mock.Mock
}

func (f FixtureServiceClient) FixtureByID(ctx context.Context, in *proto.FixtureRequest, opts ...grpc.CallOption) (*proto.Fixture, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(*proto.Fixture), args.Error(1)
}

type FixtureClient struct {
	mock.Mock
}

func (f FixtureClient) FixtureByID(id uint64) (*proto.Fixture, error) {
	args := f.Called(id)
	return args.Get(0).(*proto.Fixture), args.Error(1)
}
