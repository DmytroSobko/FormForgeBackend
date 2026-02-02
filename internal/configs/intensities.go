package configs

import (
	"fmt"

	models "github.com/DmytroSobko/FormForgeBackend/internal/models"
)

type IntensityEnvelope struct {
	Version     string                             `json:"version"`
	Intensities map[models.IntensityType]Intensity `json:"intensities"`
}

type Intensity struct {
	Multiplier        float64 `json:"multiplier"`
	FatigueMultiplier float64 `json:"fatigueMultiplier"`
}

func (i Intensity) Validate(intensity models.IntensityType) error {
	if i.Multiplier <= 0 {
		return fmt.Errorf("intensity %s: multiplier must be > 0", intensity)
	}
	if i.FatigueMultiplier <= 0 {
		return fmt.Errorf("intensity %s: fatigueMultiplier must be > 0", intensity)
	}
	return nil
}
