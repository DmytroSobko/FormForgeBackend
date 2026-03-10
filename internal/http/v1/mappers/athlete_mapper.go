package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
)

func ToAthleteResponse(a *athlete.Athlete) dto.AthleteResponse {
	return dto.AthleteResponse{
		ID:         a.GetID(),
		Type:       a.GetType(),
		Name:       a.GetName(),
		Strength:   a.GetStrength().Value(),
		Endurance:  a.GetEndurance().Value(),
		Mobility:   a.GetMobility().Value(),
		Fatigue:    a.GetFatigue(),
		MaxFatigue: a.GetMaxFatigue(),
		Week:       a.GetWeek(),
	}
}

func ToAthleteResponses(athletes []*athlete.Athlete) []dto.AthleteResponse {
	resp := make([]dto.AthleteResponse, 0, len(athletes))

	for _, a := range athletes {
		resp = append(resp, ToAthleteResponse(a))
	}

	return resp
}
