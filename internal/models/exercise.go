package models

import (
	"fmt"
)

type Exercise struct {
	ID                  string    `json:"id"`
	DisplayName         string    `json:"displayName"`
	Description         string    `json:"description"`
	PrimaryStat         StatType  `json:"primaryStat"`
	SecondaryStat       *StatType `json:"secondaryStat,omitempty"`
	SecondaryStatWeight float64   `json:"secondaryStatWeight"`
	BaseGain            float64   `json:"baseGain"`
	FatigueCost         float64   `json:"fatigueCost"`
	DurationMinutes     int       `json:"durationMinutes"`
	AllowedIntensities  []string  `json:"allowedIntensities"`
}

func (e Exercise) Validate() error {
	if e.ID == "" {
		return fmt.Errorf("exercise id is empty")
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
		return fmt.Errorf("exercise %s: secondaryStatWeight must be [0,1]", e.ID)
	}

	if e.BaseGain <= 0 {
		return fmt.Errorf("exercise %s: baseGain must be > 0", e.ID)
	}

	if e.FatigueCost < 0 {
		return fmt.Errorf("exercise %s: fatigueCost must be >= 0", e.ID)
	}

	if e.DurationMinutes <= 0 {
		return fmt.Errorf("exercise %s: durationMinutes must be > 0", e.ID)
	}

	if len(e.AllowedIntensities) == 0 {
		return fmt.Errorf("exercise %s: no allowed intensities", e.ID)
	}

	return nil
}
