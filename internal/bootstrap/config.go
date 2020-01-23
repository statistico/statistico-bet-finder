package bootstrap

import "os"

type Config struct {
	BetFair
	DataService
	OddsCompilerService
}

type BetFair struct {
	Username string
	Password string
	Key      string
}

type DataService struct {
	Host string
	Port string
}

type OddsCompilerService struct {
	Host string
	Port string
}

func BuildConfig() *Config {
	config := Config{}

	config.BetFair = BetFair{
		Username: os.Getenv("BETFAIR_USERNAME"),
		Password: os.Getenv("BETFAIR_PASSWORD"),
		Key:      os.Getenv("BETFAIR_KEY"),
	}

	config.DataService = DataService{
		Host: os.Getenv("DATA_SERVICE_HOST"),
		Port: os.Getenv("DATA_SERVICE_PORT"),
	}

	config.OddsCompilerService = OddsCompilerService{
		Host: os.Getenv("ODDS_COMPILER_SERVICE_HOST"),
		Port: os.Getenv("ODDS_COMPILER_SERVICE_PORT"),
	}

	return &config
}
