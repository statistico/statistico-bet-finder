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

type OddsCompilerServiceClient struct {
	mock.Mock
}

func (f OddsCompilerServiceClient) GetEventMarket(ctx context.Context, in *proto.EventRequest, opts ...grpc.CallOption) (*proto.EventMarket, error) {
	args := f.Called(ctx, in, opts)
	return args.Get(0).(*proto.EventMarket), args.Error(1)
}

type OddsCompilerClient struct {
	mock.Mock
}

func (o OddsCompilerClient) EventMarket(eventID uint64, market string) (*proto.EventMarket, error) {
	args := o.Called(eventID, market)
	return args.Get(0).(*proto.EventMarket), args.Error(1)
}