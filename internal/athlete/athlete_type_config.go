package athlete

import (
	"errors"
	"fmt"
)

type AthleteTypeConfig struct {
	Type               AthleteType
	DisplayName        string
	Description        string
	BaseStats          StatBlock
	MaxFatigue         float64
	RecoveryMultiplier float64
	FatigueSensitivity float64
}

type StatBlock struct {
	Strength  Stat
	Endurance Stat
	Mobility  Stat
}

func NewAthleteTypeConfig(
	athleteType string,
	displayName string,
	description string,
	strength float64,
	endurance float64,
	mobility float64,
	maxFatigue float64,
	recoveryMultiplier float64,
	fatigueSensitivity float64,
) (AthleteTypeConfig, error) {

	t := AthleteType(athleteType)

	if !t.IsValid() {
		return AthleteTypeConfig{}, fmt.Errorf("invalid athlete type: %s", athleteType)
	}

	str, err := NewStat(strength)
	if err != nil {
		return AthleteTypeConfig{}, fmt.Errorf("strength: %w", err)
	}

	end, err := NewStat(endurance)
	if err != nil {
		return AthleteTypeConfig{}, fmt.Errorf("endurance: %w", err)
	}

	mob, err := NewStat(mobility)
	if err != nil {
		return AthleteTypeConfig{}, fmt.Errorf("mobility: %w", err)
	}

	if maxFatigue <= 0 {
		return AthleteTypeConfig{}, errors.New("maxFatigue must be > 0")
	}

	if recoveryMultiplier <= 0 {
		return AthleteTypeConfig{}, errors.New("recoveryMultiplier must be > 0")
	}

	if fatigueSensitivity <= 0 {
		return AthleteTypeConfig{}, errors.New("fatigueSensitivity must be > 0")
	}

	return AthleteTypeConfig{
		Type:        t,
		DisplayName: displayName,
		Description: description,
		BaseStats: StatBlock{
			Strength:  str,
			Endurance: end,
			Mobility:  mob,
		},
		MaxFatigue:         maxFatigue,
		RecoveryMultiplier: recoveryMultiplier,
		FatigueSensitivity: fatigueSensitivity,
	}, nil
}
