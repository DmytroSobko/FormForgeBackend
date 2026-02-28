package simulation

import (
	"fmt"
)

type Intensity struct {
	Type          IntensityType
	Multiplier    float64
	FatigueFactor float64
}

func NewIntensity(
	intensityType string,
	multiplier float64,
	fatigueFactor float64,
) (Intensity, error) {

	t := IntensityType(intensityType)

	if !t.IsValid() {
		return Intensity{}, fmt.Errorf("invalid intensity type: %s", intensityType)
	}

	if multiplier <= 0 {
		return Intensity{}, fmt.Errorf("multiplier must be > 0")
	}

	if fatigueFactor < 0 {
		return Intensity{}, fmt.Errorf("fatigueFactor must be >= 0")
	}

	return Intensity{
		Type:          t,
		Multiplier:    multiplier,
		FatigueFactor: fatigueFactor,
	}, nil
}
