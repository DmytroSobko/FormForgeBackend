package app

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

type Dependencies struct {
	DB *db.DB

	Engine *simulation.Engine

	AthleteService *athlete.Service

	AthleteTypes []athlete.AthleteTypeConfig
	Exercises    []simulation.Exercise
	Intensities  []simulation.Intensity
	SimConfig    simulation.Config
}
