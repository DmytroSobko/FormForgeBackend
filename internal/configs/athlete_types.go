package configs

import "fmt"

type AthleteTypesEnvelope struct {
	Version  string        `json:"version"`
	Athletes []AthleteType `json:"athletes"`
}

type AthleteType struct {
	ID                 string    `json:"id"`
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
	if a.ID == "" {
		return fmt.Errorf("athlete id is empty")
	}

	if err := a.BaseStats.Validate(); err != nil {
		return fmt.Errorf("athlete %s: %w", a.ID, err)
	}

	if a.MaxFatigue <= 0 {
		return fmt.Errorf("athlete %s: maxFatigue must be > 0", a.ID)
	}

	if a.RecoveryMultiplier <= 0 {
		return fmt.Errorf("athlete %s: recoveryMultiplier must be > 0", a.ID)
	}

	if a.FatigueSensitivity <= 0 {
		return fmt.Errorf("athlete %s: fatigueSensitivity must be > 0", a.ID)
	}

	return nil
}

func (s StatBlock) Validate() error {
	if s.Strength < 0 || s.Endurance < 0 || s.Mobility < 0 {
		return fmt.Errorf("stats must be non-negative")
	}
	return nil
}