package configs

import (
	"fmt"

	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type ExerciseDTO struct {
	Type                string  `json:"type"`
	DisplayName         string  `json:"displayName"`
	Description         string  `json:"description"`
	PrimaryStat         string  `json:"primaryStat"`
	SecondaryStat       *string `json:"secondaryStat,omitempty"`
	SecondaryStatWeight float64 `json:"secondaryStatWeight"`
	BaseGain            float64 `json:"baseGain"`
	FatigueCost         float64 `json:"fatigueCost"`
	DurationMinutes     int     `json:"durationMinutes"`
}

type ExercisesEnvelope struct {
	Version   string        `json:"version"`
	Exercises []ExerciseDTO `json:"exercises"`
}

func LoadExercises(path string) ([]simulation.Exercise, string, error) {
	var cfg ExercisesEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, "", err
	}

	if cfg.Version == "" {
		return nil, "", fmt.Errorf("exercise config missing version")
	}

	types := map[string]bool{}
	var exercises []simulation.Exercise

	for _, dto := range cfg.Exercises {

		if types[dto.Type] {
			return nil, "", fmt.Errorf("duplicate exercise id: %s", dto.Type)
		}
		types[dto.Type] = true

		// Map DTO → Domain
		ex, err := simulation.NewExercise(
			dto.Type,
			dto.DisplayName,
			dto.Description,
			dto.PrimaryStat,
			dto.SecondaryStat,
			dto.SecondaryStatWeight,
			dto.BaseGain,
			dto.FatigueCost,
			dto.DurationMinutes,
		)
		if err != nil {
			return nil, "", err
		}

		exercises = append(exercises, ex)
	}

	return exercises, cfg.Version, nil
}
