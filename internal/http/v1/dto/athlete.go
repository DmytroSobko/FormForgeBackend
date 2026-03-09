package dto

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
)

type CreateAthleteRequest struct {
	Type athlete.AthleteType `json:"type" validate:"required"`
	Name string              `json:"name" validate:"required,min=2,max=50"`
}

type AthleteResponse struct {
	ID         string              `json:"id"`
	Type       athlete.AthleteType `json:"type"`
	Name       string              `json:"name"`
	Strength   float64             `json:"strength"`
	Endurance  float64             `json:"endurance"`
	Mobility   float64             `json:"mobility"`
	Fatigue    float64             `json:"fatigue"`
	MaxFatigue float64             `json:"maxFatigue"`
	Week       int                 `json:"week"`
}
