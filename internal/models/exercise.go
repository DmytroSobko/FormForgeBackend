package models

import (
	"fmt"
)

type Exercise struct {
	Type                string    `json:"type"`
	DisplayName         string    `json:"displayName"`
	Description         string    `json:"description"`
	PrimaryStat         StatType  `json:"primaryStat"`
	SecondaryStat       *StatType `json:"secondaryStat,omitempty"`
	SecondaryStatWeight float64   `json:"secondaryStatWeight"`
	BaseGain            float64   `json:"baseGain"`
	FatigueCost         float64   `json:"fatigueCost"`
	DurationMinutes     int       `json:"durationMinutes"`
}

func (e Exercise) Validate() error {
	if e.Type == "" {
		return fmt.Errorf("exercise type is empty")
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

	if e.SecondaryStatWeight < 0 || e.SecondaryStatWeight > 1 {
		return fmt.Errorf("exercise %s: secondaryStatWeight must be [0,1]", e.Type)
	}

	if e.BaseGain <= 0 {
		return fmt.Errorf("exercise %s: baseGain must be > 0", e.Type)
	}

	if e.FatigueCost < 0 {
		return fmt.Errorf("exercise %s: fatigueCost must be >= 0", e.Type)
	}

	if e.DurationMinutes <= 0 {
		return fmt.Errorf("exercise %s: durationMinutes must be > 0", e.Type)
	}

	return nil
}
