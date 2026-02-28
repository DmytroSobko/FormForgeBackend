package configs

import (
	"fmt"

	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
)

type AthleteTypeDTO struct {
	Type               string       `json:"type"`
	DisplayName        string       `json:"displayName"`
	Description        string       `json:"description"`
	BaseStats          StatBlockDTO `json:"baseStats"`
	MaxFatigue         float64      `json:"maxFatigue"`
	RecoveryMultiplier float64      `json:"recoveryMultiplier"`
	FatigueSensitivity float64      `json:"fatigueSensitivity"`
}

type StatBlockDTO struct {
	Strength  float64 `json:"strength"`
	Endurance float64 `json:"endurance"`
	Mobility  float64 `json:"mobility"`
}

type AthleteTypesEnvelope struct {
	Version      string           `json:"version"`
	AthleteTypes []AthleteTypeDTO `json:"athleteTypes"`
}

func LoadAthleteTypes(path string) ([]athlete.AthleteTypeConfig, string, error) {
	var cfg AthleteTypesEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, "", err
	}

	if cfg.Version == "" {
		return nil, "", fmt.Errorf("athlete type config missing version")
	}

	types := map[string]bool{}
	var result []athlete.AthleteTypeConfig

	for _, dto := range cfg.AthleteTypes {

		if types[dto.Type] {
			return nil, "", fmt.Errorf("duplicate athlete type: %s", dto.Type)
		}
		types[dto.Type] = true

		at, err := athlete.NewAthleteTypeConfig(
			dto.Type,
			dto.DisplayName,
			dto.Description,
			dto.BaseStats.Strength,
			dto.BaseStats.Endurance,
			dto.BaseStats.Mobility,
			dto.MaxFatigue,
			dto.RecoveryMultiplier,
			dto.FatigueSensitivity,
		)
		if err != nil {
			return nil, "", err
		}

		result = append(result, at)
	}

	return result, cfg.Version, nil
}
