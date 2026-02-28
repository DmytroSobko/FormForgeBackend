package simulation

import "errors"

type Config struct {
	RestDayRecovery      float64
	MaxFatiguePenalty    float64
	HighFatigueThreshold float64
}

func NewConfig(
	restDayRecovery float64,
	maxFatiguePenalty float64,
	highFatigueThreshold float64,
) (Config, error) {

	if restDayRecovery < 0 {
		return Config{}, errors.New("restDayRecovery must be >= 0")
	}

	if maxFatiguePenalty <= 0 || maxFatiguePenalty > 1 {
		return Config{}, errors.New("maxFatiguePenalty must be between 0 and 1")
	}

	if highFatigueThreshold < 0 || highFatigueThreshold > maxFatiguePenalty {
		return Config{}, errors.New("highFatigueThreshold must be between 0 and maxFatiguePenalty")
	}

	return Config{
		RestDayRecovery:      restDayRecovery,
		MaxFatiguePenalty:    maxFatiguePenalty,
		HighFatigueThreshold: highFatigueThreshold,
	}, nil
}
