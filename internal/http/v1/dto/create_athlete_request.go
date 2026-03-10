package dto

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
)

type CreateAthleteRequest struct {
	Type athlete.AthleteType `json:"type" validate:"required"`
	Name string              `json:"name" validate:"required,min=2,max=50"`
}
