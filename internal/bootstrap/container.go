package bootstrap

import (
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-bet-finder/internal/app/grpc/proto"
	bfClient "github.com/statistico/statistico-betfair-go-client"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"time"
)

type Container struct {
	BetFairClient    *bfClient.Client
	Clock            clockwork.Clock
	Config           *Config
	FixtureClient    proto.FixtureServiceClient
	Logger           *logrus.Logger
	OddsCompilerClient proto.OddsCompilerServiceClient
}

func BuildContainer(config *Config) *Container {
	c := Container{
		Config: config,
	}

	c.BetFairClient = betFairClient(config)
	c.Clock = clock()
	c.FixtureClient = fixtureClient(config)
	c.Logger = logger()
	c.OddsCompilerClient = oddsCompilerClient(config)

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
		HTTPClient:  client,
		Credentials: bfClient.InteractiveCredentials{
			Username: config.BetFair.Username,
			Password: config.BetFair.Password,
			Key: config.BetFair.Key,
		},
		BaseURLs:    bfClient.BaseURLs{
			Accounts:  "https://api.betfair.com/exchange/account/rest/v1.0/",
			Betting:   "https://api.betfair.com/exchange/betting/rest/v1.0/",
			Login:     "https://identitysso.betfair.com/api/login",
		},
	}

	return &betfair
}

func fixtureClient(config *Config) proto.FixtureServiceClient {
	conn, err := grpc.Dial(config.DataService.Host + ":" + config.DataService.Port, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	return proto.NewFixtureServiceClient(conn)
}

func logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	return logger
}

func oddsCompilerClient(config *Config) proto.OddsCompilerServiceClient {
	conn, err := grpc.Dial(config.OddsCompilerService.Host + ":" + config.OddsCompilerService.Port, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	return proto.NewOddsCompilerServiceClient(conn)
}

func clock() clockwork.Clock {
	return clockwork.NewRealClock()
}
