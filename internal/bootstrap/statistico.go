package bootstrap

import "github.com/statistico/statistico-bet-finder/internal/app/statistico"

func (c Container) StatisticoFixtureClient() statistico.FixtureClient {
	return statistico.NewGRPCFixtureClient(c.FixtureClient)
}

func (c Container) StatisticoOddsCompilerClient() statistico.OddsCompilerClient {
	return statistico.NewGRPCOddsCompilerClient(c.OddsCompilerClient)
}
