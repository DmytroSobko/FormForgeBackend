package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func ToIntensityConfig(i simulation.Intensity) dto.IntensityConfig {
	return dto.IntensityConfig{
		Type:          i.Type.String(),
		Multiplier:    i.Multiplier,
		FatigueFactor: i.FatigueFactor,
	}
}
