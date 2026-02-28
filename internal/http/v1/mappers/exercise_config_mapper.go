package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func ToExerciseConfig(e simulation.Exercise) dto.ExerciseConfig {

	var secondary *string
	if e.SecondaryStat != nil {
		s := e.SecondaryStat.String()
		secondary = &s
	}

	return dto.ExerciseConfig{
		Type:                e.Type,
		DisplayName:         e.DisplayName,
		Description:         e.Description,
		PrimaryStat:         e.PrimaryStat.String(),
		SecondaryStat:       secondary,
		SecondaryStatWeight: e.SecondaryStatWeight,
		BaseGain:            e.BaseGain,
		FatigueCost:         e.FatigueCost,
		DurationMinutes:     e.DurationMinutes,
	}
}
