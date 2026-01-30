package models

import (
	"errors"
	"fmt"
)

type StatType string

const (
	StatStrength  StatType = "strength"
	StatEndurance StatType = "endurance"
	StatMobility  StatType = "mobility"
)

type IntensityType string

const (
	IntensityLow    IntensityType = "low"
	IntensityMedium IntensityType = "medium"
	IntensityHigh   IntensityType = "high"
)

type Exercise struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	PrimaryStat     StatType      `json:"primaryStat"`
	SecondaryStat   *StatType     `json:"secondaryStat,omitempty"`
	Intensity       IntensityType `json:"intensity"`
	BaseGain        float64       `json:"baseGain"`
	FatigueCost     float64       `json:"fatigueCost"`
	DurationMinutes int           `json:"durationMinutes"`
}

func (e Exercise) Validate() error {
	if e.ID == "" {
		return errors.New("exercise id cannot be empty")
	}

	if e.Name == "" {
		return errors.New("exercise name cannot be empty")
	}

	if e.BaseGain < 0 {
		return errors.New("baseGain must be >= 0")
	}

	if e.FatigueCost < 0 {
		return errors.New("fatigueCost must be >= 0")
	}

	if e.DurationMinutes <= 0 {
		return errors.New("durationMinutes must be > 0")
	}

	switch e.PrimaryStat {
	case StatStrength, StatEndurance, StatMobility:
		// ok
	default:
		return fmt.Errorf("invalid primaryStat: %s", e.PrimaryStat)
	}

	if e.SecondaryStat != nil {
		switch *e.SecondaryStat {
		case StatStrength, StatEndurance, StatMobility:
			// ok
		default:
			return fmt.Errorf("invalid secondaryStat: %s", *e.SecondaryStat)
		}
	}

	switch e.Intensity {
	case IntensityLow, IntensityMedium, IntensityHigh:
		// ok
	default:
		return fmt.Errorf("invalid intensity: %s", e.Intensity)
	}

	return nil
}
