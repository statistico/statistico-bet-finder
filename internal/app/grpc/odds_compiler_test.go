package grpc_test

import (
	"context"
	"errors"
	grpc2 "github.com/statistico/statistico-price-finder/internal/app/grpc"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"github.com/statistico/statistico-price-finder/internal/app/mock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestOddsCompilerClient_EventAndMarket(t *testing.T) {
	t.Run("returns event market struct", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.OddsCompilerServiceClient)
		oddsClient := grpc2.NewOddsCompilerClient(mockClient)
		request := proto.EventRequest{
			EventId:              45019,
			Market:               "OVER_UNDER_25",
		}

		market := proto.EventMarket{
			EventId:              45019,
			Market:               "OVER_UNDER_25",
			Odds:                 []*proto.Odds{
				{
					Price: 1.96,
					Selection: "Over 2.5 Goals",
				},
				{
					Price: 2.54,
					Selection: "Under 2.5 Goals",
				},
			},
		}

		mockClient.On("GetEventMarket", context.Background(), &request, []grpc.CallOption(nil)).Return(&market, nil)

		em, err := oddsClient.EventMarket(45019, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Error calling fixture client expected nil got %s", err)
		}

		mockClient.AssertExpectations(t)
		assert.Equal(t, uint64(45019), em.EventId)
		assert.Equal(t, "OVER_UNDER_25", em.Market)
		assert.Equal(t, "Over 2.5 Goals", em.Odds[0].Selection)
		assert.Equal(t, float32(1.96), em.Odds[0].Price)
		assert.Equal(t, "Under 2.5 Goals", em.Odds[1].Selection)
		assert.Equal(t, float32(2.54), em.Odds[1].Price)
	})

	t.Run("returns error if error returns by odds compiler client", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.OddsCompilerServiceClient)
		oddsClient := grpc2.NewOddsCompilerClient(mockClient)
		request := proto.EventRequest{
			EventId:              45019,
			Market:               "OVER_UNDER_25",
		}

		mockClient.On("GetEventMarket", context.Background(), &request, []grpc.CallOption(nil)).Return(&proto.EventMarket{}, errors.New("client error"))

		em, err := oddsClient.EventMarket(45019, "OVER_UNDER_25")

		if em != nil {
			t.Fatalf("Error calling odds compiler expected nil got %s", err)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(t, "client error", err.Error())
	})
}
