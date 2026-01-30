package configs

import "fmt"

type ExerciseEnvelope struct {
	Version   string     `json:"version"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	ID                  string   `json:"id"`
	DisplayName         string   `json:"displayName"`
	Description         string   `json:"description"`
	PrimaryStat         string   `json:"primaryStat"`
	SecondaryStat       *string  `json:"secondaryStat"`
	SecondaryStatWeight float64  `json:"secondaryStatWeight"`
	BaseGain            float64  `json:"baseGain"`
	FatigueCost         float64  `json:"fatigueCost"`
	DurationMinutes     int      `json:"durationMinutes"`
	AllowedIntensities  []string `json:"allowedIntensities"`
}

func (e Exercise) Validate() error {
	if e.ID == "" {
		return fmt.Errorf("exercise id is empty")
	}

	if !isValidStat(e.PrimaryStat) {
		return fmt.Errorf("exercise %s: invalid primaryStat", e.ID)
	}

	if e.SecondaryStat != nil && !isValidStat(*e.SecondaryStat) {
		return fmt.Errorf("exercise %s: invalid secondaryStat", e.ID)
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

func isValidStat(stat string) bool {
	switch stat {
	case "strength", "endurance", "mobility":
		return true
	default:
		return false
	}
}
