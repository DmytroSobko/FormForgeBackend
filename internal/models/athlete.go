package models

import "errors"

type Athlete struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Strength   float64 `json:"strength"`
	Endurance  float64 `json:"endurance"`
	Mobility   float64 `json:"mobility"`
	Fatigue    float64 `json:"fatigue"`
	MaxFatigue float64 `json:"maxFatigue"`
	Week       int     `json:"week"`
}

func NewAthlete(id string, name string, strength float64, endurance float64, mobility float64, maxFatigue float64) Athlete {
	return Athlete{
		ID:         id,
		Name:       name,
		Strength:   strength,
		Endurance:  endurance,
		Mobility:   mobility,
		Fatigue:    0,
		MaxFatigue: maxFatigue,
		Week:       1,
	}
}

func (a Athlete) Validate() error {
	if a.ID == "" {
		return errors.New("athlete id cannot be empty")
	}

	if a.Name == "" {
		return errors.New("athlete name cannot be empty")
	}

	if err := validateStatRange("strength", a.Strength); err != nil {
		return err
	}
	if err := validateStatRange("endurance", a.Endurance); err != nil {
		return err
	}
	if err := validateStatRange("mobility", a.Mobility); err != nil {
		return err
	}

	if err := validateNonNegative("fatigue", a.Fatigue); err != nil {
		return err
	}

	if a.MaxFatigue <= 0 {
		return errors.New("maxFatigue must be greater than 0")
	}

	if a.Fatigue > a.MaxFatigue {
		return errors.New("fatigue cannot exceed maxFatigue")
	}

	if a.Week < 1 {
		return errors.New("week must be >= 1")
	}

	return nil
}
