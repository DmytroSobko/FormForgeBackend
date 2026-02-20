package configs

import "errors"

type SimulationConfigEnvelope struct {
	Version    string           `json:"version"`
	Simulation SimulationConfig `json:"simulation"`
}

type SimulationConfig struct {
	RestDayRecovery      float64 `json:"restDayRecovery"`
	MaxFatiguePenalty    float64 `json:"maxFatiguePenalty"`
	HighFatigueThreshold float64 `json:"highFatigueThreshold"`
}

func (c SimulationConfig) Validate() error {
	if c.RestDayRecovery < 0 {
		return errors.New("restDayRecovery must be >= 0")
	}
	if c.MaxFatiguePenalty <= 0 || c.MaxFatiguePenalty > 1 {
		return errors.New("maxFatiguePenalty must be between 0 and 1")
	}
	if c.HighFatigueThreshold < 0 || c.HighFatigueThreshold > c.MaxFatiguePenalty {
		return errors.New("highFatigueThreshold must be between 0 and maxFatiguePenalty")
	}

	return nil
}
