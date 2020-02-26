package bootstrap

import (
	"fmt"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-price-finder/internal/app/grpc/proto"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"time"
)

type Container struct {
	BetFairClient *bfClient.Client
	Clock         clockwork.Clock
	Config        *Config
	FixtureClient proto.FixtureServiceClient
	Logger        *logrus.Logger
}

func BuildContainer(config *Config) *Container {
	c := Container{
		Config: config,
	}

	c.BetFairClient = betFairClient(config)
	c.Clock = clock()
	c.FixtureClient = fixtureClient(config)
	c.Logger = logger()

	return &c
}

func betFairClient(config *Config) *bfClient.Client {
	trans := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 30,
		Transport: trans,
	}

	betfair := bfClient.Client{
		HTTPClient: client,
		Credentials: bfClient.InteractiveCredentials{
			Username: config.BetFair.Username,
			Password: config.BetFair.Password,
			Key:      config.BetFair.Key,
		},
		BaseURLs: bfClient.BaseURLs{
			Accounts: bfClient.AccountsURL,
			Betting:  bfClient.BettingURL,
			Login:    bfClient.LoginURL,
		},
	}

	return &betfair
}

func fixtureClient(config *Config) proto.FixtureServiceClient {
	host := fmt.Sprintf("%s:%s", config.DataService.Host, config.DataService.Port)

	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	return proto.NewFixtureServiceClient(conn)
}

func logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func clock() clockwork.Clock {
	return clockwork.NewRealClock()
}
