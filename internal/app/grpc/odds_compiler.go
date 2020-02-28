package grpc

import (
	"context"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
)

type OddsCompilerClient interface {
	EventMarket(eventID uint64, market string) (*proto.EventMarket, error)
}

// oddsCompilerClient is a wrapper around the Statistico Odds Compiler service.
type oddsCompilerClient struct {
	client proto.OddsCompilerServiceClient
}

// EventMarket returns an EventMarket struct parsed from the Statistico Odds Compiler service.
func (o oddsCompilerClient) EventMarket(eventID uint64, market string) (*proto.EventMarket, error) {
	request := proto.EventRequest{
		EventId: eventID,
		Market:  market,
	}

	response, err := o.client.GetEventMarket(context.Background(), &request)

	if err != nil {
		return nil, err
	}

	return response, err
}

func NewOddsCompilerClient(c proto.OddsCompilerServiceClient) OddsCompilerClient {
	return &oddsCompilerClient{client: c}
}
