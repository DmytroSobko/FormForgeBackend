package configs

import (
	models "github.com/DmytroSobko/FormForgeBackend/internal/models"
)

type ExercisesEnvelope struct {
	Version   string            `json:"version"`
	Exercises []models.Exercise `json:"exercises"`
}
