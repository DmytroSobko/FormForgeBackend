package configs

import "fmt"

type AthleteTypesEnvelope struct {
	Version       string        `json:"version"`
	AthletesTypes []AthleteType `json:"athleteTypes"`
}

type AthleteType struct {
	Type               string    `json:"type"`
	DisplayName        string    `json:"displayName"`
	Description        string    `json:"description"`
	BaseStats          StatBlock `json:"baseStats"`
	MaxFatigue         float64   `json:"maxFatigue"`
	RecoveryMultiplier float64   `json:"recoveryMultiplier"`
	FatigueSensitivity float64   `json:"fatigueSensitivity"`
}

type StatBlock struct {
	Strength  float64 `json:"strength"`
	Endurance float64 `json:"endurance"`
	Mobility  float64 `json:"mobility"`
}

func (a AthleteType) Validate() error {
	if a.Type == "" {
		return fmt.Errorf("athlete type is empty")
	}

	if err := a.BaseStats.Validate(); err != nil {
		return fmt.Errorf("athlete %s: %w", a.Type, err)
	}

	if a.MaxFatigue <= 0 {
		return fmt.Errorf("athlete %s: maxFatigue must be > 0", a.Type)
	}

	if a.RecoveryMultiplier <= 0 {
		return fmt.Errorf("athlete %s: recoveryMultiplier must be > 0", a.Type)
	}

	if a.FatigueSensitivity <= 0 {
		return fmt.Errorf("athlete %s: fatigueSensitivity must be > 0", a.Type)
	}

	return nil
}

func (s StatBlock) Validate() error {
	if s.Strength < 0 || s.Endurance < 0 || s.Mobility < 0 {
		return fmt.Errorf("stats must be non-negative")
	}
	return nil
}
