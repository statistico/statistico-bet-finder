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

func TestFixtureClient_FixtureByID(t *testing.T) {
	t.Run("returns fixture struct", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.FixtureServiceClient)
		fixtureClient := FixtureClient{client:mockClient}
		request := proto.FixtureRequest{FixtureId:14562}

		fix := proto.Fixture{
			Id:                   14562,
			Competition:          &proto.Competition{Id:42},
			HomeTeam:             &proto.Team{Name:"West Ham United"},
			AwayTeam:             &proto.Team{Name:"Arsenal"},
			DateTime:             &proto.Date{Utc:1579536616},
		}

		mockClient.On("FixtureByID", context.Background(), &request, []grpc.CallOption(nil)).Return(&fix, nil)

		fixture, err := fixtureClient.FixtureByID(14562)

		if err != nil {
			t.Fatalf("Error calling fixture client expected nil got %s", err)
		}

		mockClient.AssertExpectations(t)
		assert.Equal(t, uint64(14562), fixture.ID)
		assert.Equal(t, uint64(42), fixture.CompetitionID)
		assert.Equal(t, "West Ham United", fixture.HomeTeam)
		assert.Equal(t, "Arsenal", fixture.AwayTeam)
		assert.Equal(t, int64(1579536616), fixture.Date.UTC().Unix())
	})

	t.Run("returns error if error returned calling fixture client", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.FixtureServiceClient)
		fixtureClient := FixtureClient{client:mockClient}
		request := proto.FixtureRequest{FixtureId:14562}

		mockClient.On("FixtureByID", context.Background(), &request, []grpc.CallOption(nil)).Return(&proto.Fixture{}, errors.New("client error"))

		fixture, err := fixtureClient.FixtureByID(14562)

		if fixture != nil {
			t.Fatalf("Error calling odds compiler expected nil got %s", err)
		}

		if err == nil {
			t.Fatal("Error expected error got nil")
		}

		assert.Equal(t, "client error", err.Error())
	})
}
