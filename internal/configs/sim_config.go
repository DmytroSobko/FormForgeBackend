package configs

import "errors"

type SimulationConfigEnvelope struct {
	Version    string           `json:"version"`
	Simulation SimulationConfig `json:"simulation"`
}

type SimulationConfig struct {
	DaysInWeek           int                `json:"daysInWeek"`
	RestDayRecovery      float64            `json:"restDayRecovery"`
	MaxFatiguePenalty    float64            `json:"maxFatiguePenalty"`
	HighFatigueThreshold float64            `json:"highFatigueThreshold"`
	IntensityMultipliers map[string]float64 `json:"intensityMultipliers"`
}

func (c SimulationConfig) Validate() error {
	if c.DaysInWeek <= 0 {
		return errors.New("daysInWeek must be > 0")
	}
	if c.RestDayRecovery < 0 {
		return errors.New("restDayRecovery must be >= 0")
	}
	if c.MaxFatiguePenalty <= 0 || c.MaxFatiguePenalty > 1 {
		return errors.New("maxFatiguePenalty must be between 0 and 1")
	}
	if c.HighFatigueThreshold < 0 || c.HighFatigueThreshold > c.MaxFatiguePenalty {
		return errors.New("highFatigueThreshold must be between 0 and maxFatiguePenalty")
	}

	required := []string{"low", "medium", "high"}
	for _, key := range required {
		if _, ok := c.IntensityMultipliers[key]; !ok {
			return errors.New("missing intensity multiplier: " + key)
		}
	}

	return nil
}
