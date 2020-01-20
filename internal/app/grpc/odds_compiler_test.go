package grpc

import (
	"context"
	"errors"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
	"github.com/statistico/statistico-bet-finder/internal/app/mock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestOddsCompilerClient_GetOverUnderGoalsForFixture(t *testing.T) {
	t.Run("returns market struct containing from odds compiler client response", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.OddsCompilerServiceClient)
		oddsClient := NewOddsCompilerClient(mockClient)
		request := proto.OverUnderRequest{Market:"OVER_UNDER_15",FixtureId:145261}

		oddsOne := proto.Odds{Price: 1.67, Selection: "OVER"}
		oddsTwo := proto.Odds{Price: 4.62, Selection: "UNDER"}

		response := proto.OverUnderGoalsResponse{
			FixtureId:            145261,
			Market:               "OVER_UNDER_15",
		}

		response.Odds = append(response.Odds, &oddsOne)
		response.Odds = append(response.Odds, &oddsTwo)

		mockClient.On("GetOverUnderGoalsForFixture", context.Background(), &request, []grpc.CallOption(nil)).Return(&response, nil)

		market, err := oddsClient.GetOverUnderGoalsForFixture(145261, "OVER_UNDER_15")

		if err != nil {
			t.Fatalf("Error calling odds compiler expected nil got %s", err)
		}

		mockClient.AssertExpectations(t)
		assert.Equal(t, uint64(145261), market.FixtureID)
		assert.Equal(t, "OVER_UNDER_15", market.Name)
		assert.Equal(t, float32(1.67), market.Runners[0].Price)
		assert.Equal(t, "OVER", market.Runners[0].Name)
		assert.Equal(t, float32(4.62), market.Runners[1].Price)
		assert.Equal(t, "UNDER", market.Runners[1].Name)
	})

	t.Run("returns error if error returned calling calling odds client", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.OddsCompilerServiceClient)
		oddsClient := NewOddsCompilerClient(mockClient)
		request := proto.OverUnderRequest{Market:"OVER_UNDER_15",FixtureId:145261}

		mockClient.On("GetOverUnderGoalsForFixture", context.Background(), &request, []grpc.CallOption(nil)).Return(
			&proto.OverUnderGoalsResponse{},
			errors.New("client error"),
		)

		market, err := oddsClient.GetOverUnderGoalsForFixture(145261, "OVER_UNDER_15")

		if market != nil {
			t.Fatalf("Error calling odds compiler expected nil got %s", err)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(t, "client error", err.Error())
	})
}
