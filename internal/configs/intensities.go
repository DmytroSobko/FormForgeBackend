package configs

import "fmt"

type IntensityEnvelope struct {
	Version     string               `json:"version"`
	Intensities map[string]Intensity `json:"intensities"`
}

type Intensity struct {
	Multiplier        float64 `json:"multiplier"`
	FatigueMultiplier float64 `json:"fatigueMultiplier"`
}

func (i Intensity) Validate(name string) error {
	if i.Multiplier <= 0 {
		return fmt.Errorf("intensity %s: multiplier must be > 0", name)
	}
	if i.FatigueMultiplier <= 0 {
		return fmt.Errorf("intensity %s: fatigueMultiplier must be > 0", name)
	}
	return nil
}
