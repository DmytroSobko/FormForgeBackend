package configs

import (
	"fmt"
)

type IntensitiesEnvelope struct {
	Version     string      `json:"version"`
	Intensities []Intensity `json:"intensities"`
}

type Intensity struct {
	Type              string  `json:"type"`
	Multiplier        float64 `json:"multiplier"`
	FatigueMultiplier float64 `json:"fatigueMultiplier"`
}

func (i Intensity) Validate() error {
	if i.Type == "" {
		return fmt.Errorf("intensity type is empty")
	}
	if i.Multiplier <= 0 {
		return fmt.Errorf("intensity %s: multiplier must be > 0", i.Type)
	}
	if i.FatigueMultiplier <= 0 {
		return fmt.Errorf("intensity %s: fatigueMultiplier must be > 0", i.Type)
	}
	return nil
}
