package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
)

// -----------------------------
// AthleteTypeConfig → AthleteTypeResponse
// -----------------------------

func ToAthleteTypeConfig(t athlete.AthleteTypeConfig) dto.AthleteTypeConfig {
	return dto.AthleteTypeConfig{
		Type:        t.Type.String(),
		DisplayName: t.DisplayName,
		Description: t.Description,
		BaseStats: dto.StatBlock{
			Strength:  t.BaseStats.Strength.Value(),
			Endurance: t.BaseStats.Endurance.Value(),
			Mobility:  t.BaseStats.Mobility.Value(),
		},
		MaxFatigue:         t.MaxFatigue,
		RecoveryMultiplier: t.RecoveryMultiplier,
		FatigueSensitivity: t.FatigueSensitivity,
	}
}
