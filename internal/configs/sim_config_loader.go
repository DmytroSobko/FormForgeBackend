package configs

import (
	"fmt"

	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type SimulationConfigDTO struct {
	RestDayRecovery      float64 `json:"restDayRecovery"`
	MaxFatiguePenalty    float64 `json:"maxFatiguePenalty"`
	HighFatigueThreshold float64 `json:"highFatigueThreshold"`
}

type SimulationConfigEnvelope struct {
	Version    string              `json:"version"`
	Simulation SimulationConfigDTO `json:"simulation"`
}

func LoadSimulationConfig(path string) (simulation.Config, string, error) {
	var cfg SimulationConfigEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return simulation.Config{}, "", err
	}

	if cfg.Version == "" {
		return simulation.Config{}, "", fmt.Errorf("simulation config missing version")
	}

	domainCfg, err := simulation.NewConfig(
		cfg.Simulation.RestDayRecovery,
		cfg.Simulation.MaxFatiguePenalty,
		cfg.Simulation.HighFatigueThreshold,
	)
	if err != nil {
		return simulation.Config{}, "", err
	}

	return domainCfg, cfg.Version, nil
}
