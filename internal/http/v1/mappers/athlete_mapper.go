package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
)

// -----------------------------
// Athlete → AthleteResponse
// -----------------------------

func ToAthleteResponse(a *athlete.Athlete) dto.AthleteResponse {
	return dto.AthleteResponse{
		ID:         a.GetID(),
		Type:       a.GetType().String(),
		Name:       a.GetName(),
		Strength:   a.GetStrength().Value(),
		Endurance:  a.GetEndurance().Value(),
		Mobility:   a.GetMobility().Value(),
		Fatigue:    a.GetFatigue(),
		MaxFatigue: a.GetMaxFatigue(),
		Week:       a.GetWeek(),
	}
}
