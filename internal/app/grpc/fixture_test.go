package grpc_test

import (
	"context"
	"errors"
	grpc2 "github.com/statistico/statistico-bet-finder/internal/app/grpc"
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
		fixtureClient := grpc2.NewFixtureClient(mockClient)
		request := proto.FixtureRequest{FixtureId: 14562}

		fix := proto.Fixture{
			Id:          14562,
			Competition: &proto.Competition{Id: 42},
			HomeTeam:    &proto.Team{Name: "West Ham United"},
			AwayTeam:    &proto.Team{Name: "Arsenal"},
			DateTime:    &proto.Date{Utc: 1579536616},
		}

		mockClient.On("FixtureByID", context.Background(), &request, []grpc.CallOption(nil)).Return(&fix, nil)

		fixture, err := fixtureClient.FixtureByID(14562)

		if err != nil {
			t.Fatalf("Error calling fixture client expected nil got %s", err)
		}

		mockClient.AssertExpectations(t)
		assert.Equal(t, int64(14562), fixture.Id)
		assert.Equal(t, int64(42), fixture.Competition.Id)
		assert.Equal(t, "West Ham United", fixture.HomeTeam.Name)
		assert.Equal(t, "Arsenal", fixture.AwayTeam.Name)
		assert.Equal(t, int64(1579536616), fixture.DateTime.Utc)
	})

	t.Run("returns error if error returned calling fixture client", func(t *testing.T) {
		t.Helper()

		mockClient := new(mock.FixtureServiceClient)
		fixtureClient := grpc2.NewFixtureClient(mockClient)
		request := proto.FixtureRequest{FixtureId: 14562}

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
