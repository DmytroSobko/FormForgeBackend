package simulation

import (
	"math"

	"github.com/DmytroSobko/FormForgeBackend/internal/config"
	"github.com/DmytroSobko/FormForgeBackend/internal/models"
)

type Engine struct {
	cfg *config.SimulationConfig
}

func NewEngine(cfg *config.SimulationConfig) *Engine {
	return &Engine{
		cfg: cfg,
	}
}

func (e *Engine) SimulateWeek(
	athlete models.Athlete,
	plan models.TrainingPlan,
) models.SimulationResult {

	before := snapshot(athlete)

	totalPotential := 0.0
	totalActual := 0.0
	warnings := []string{}

	for _, day := range plan.Days {
		if len(day.Exercises) == 0 {
			athlete.Fatigue = math.Max(
				athlete.Fatigue-e.cfg.RestDayRecovery,
				0,
			)
			continue
		}

		for _, ex := range day.Exercises {
			intensity := e.cfg.IntensityMultipliers[string(ex.Intensity)]

			penalty := math.Min(
				athlete.Fatigue/athlete.MaxFatigue,
				e.cfg.MaxFatiguePenalty,
			)

			rawGain := ex.BaseGain * intensity
			finalGain := rawGain * (1 - penalty)

			applyStat(&athlete, ex.PrimaryStat, finalGain)

			athlete.Fatigue += ex.FatigueCost * intensity
			athlete.Fatigue = math.Min(
				athlete.Fatigue,
				athlete.MaxFatigue,
			)

			totalPotential += rawGain
			totalActual += finalGain

			if penalty > e.cfg.HighFatigueThreshold {
				warnings = append(warnings, "High fatigue reduced gains")
			}
		}
	}

	after := snapshot(athlete)

	efficiency := 1.0
	if totalPotential > 0 {
		efficiency = totalActual / totalPotential
	}

	return models.NewSimulationResult(
		"",
		athlete.ID,
		athlete.Week,
		before,
		after,
		efficiency,
		warnings,
	)
}

func snapshot(a models.Athlete) models.StatSnapshot {
	return models.StatSnapshot{
		Strength:  a.Strength,
		Endurance: a.Endurance,
		Mobility:  a.Mobility,
		Fatigue:   a.Fatigue,
	}
}

func applyStat(
	athlete *models.Athlete,
	stat models.StatType,
	value float64,
) {
	switch stat {
	case models.StatStrength:
		athlete.Strength += value
	case models.StatEndurance:
		athlete.Endurance += value
	case models.StatMobility:
		athlete.Mobility += value
	}
}
