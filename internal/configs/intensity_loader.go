package configs

import (
	"fmt"

	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type IntensityDTO struct {
	Type          string  `json:"type"`
	Multiplier    float64 `json:"multiplier"`
	FatigueFactor float64 `json:"fatigueFactor"`
}

type IntensitiesEnvelope struct {
	Version     string         `json:"version"`
	Intensities []IntensityDTO `json:"intensities"`
}

func LoadIntensities(path string) ([]simulation.Intensity, string, error) {
	var cfg IntensitiesEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, "", err
	}

	if cfg.Version == "" {
		return nil, "", fmt.Errorf("intensity config missing version")
	}

	types := map[string]bool{}
	var intensities []simulation.Intensity

	for _, dto := range cfg.Intensities {

		if types[dto.Type] {
			return nil, "", fmt.Errorf("duplicate intensity type: %s", dto.Type)
		}
		types[dto.Type] = true

		intensity, err := simulation.NewIntensity(
			dto.Type,
			dto.Multiplier,
			dto.FatigueFactor,
		)
		if err != nil {
			return nil, "", err
		}

		intensities = append(intensities, intensity)
	}

	return intensities, cfg.Version, nil
}
