package configs

import (
	models "github.com/DmytroSobko/FormForgeBackend/internal/models"
)

type ExerciseEnvelope struct {
	Version   string            `json:"version"`
	Exercises []models.Exercise `json:"exercises"`
}
