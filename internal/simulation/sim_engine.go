package simulation

import (
	"fmt"
	"math"

	"github.com/DmytroSobko/FormForgeBackend/internal/configs"
	"github.com/DmytroSobko/FormForgeBackend/internal/models"
)

type Engine struct {
	cfg         *configs.SimulationConfig
	intensities []configs.Intensity
}

func NewEngine(cfg *configs.SimulationConfig,
	intensities []configs.Intensity) *Engine {
	return &Engine{
		cfg:         cfg,
		intensities: intensities,
	}
}

func (e *Engine) SimulateWeek(athlete models.Athlete, plan models.TrainingPlan) (*models.SimulationResult, error) {

	before := snapshot(athlete)

	totalPotential := 0.0
	totalActual := 0.0
	warnings := []string{}

	for _, day := range plan.Days {

		// Rest day
		if len(day.Exercises) == 0 {
			athlete.Fatigue = math.Max(
				athlete.Fatigue-e.cfg.RestDayRecovery,
				0,
			)
			continue
		}

		for _, planned := range day.Exercises {

			ex := planned.Exercise

			intensityCfg, err := getIntensityConfig(planned.Intensity, e.intensities)

			if err != nil {
				return nil, fmt.Errorf("intensity %s not found", planned.Intensity)
			}

			fatigueRatio := athlete.Fatigue / athlete.MaxFatigue
			penalty := math.Min(fatigueRatio, e.cfg.MaxFatiguePenalty)

			rawGain := ex.BaseGain * intensityCfg.Multiplier
			finalGain := rawGain * (1 - penalty)

			applyStat(&athlete, ex.PrimaryStat, finalGain)

			athlete.Fatigue += ex.FatigueCost * intensityCfg.FatigueMultiplier
			athlete.Fatigue = math.Min(athlete.Fatigue, athlete.MaxFatigue)

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

	simRes := models.NewSimulationResult("", athlete.ID, athlete.Week, before, after, efficiency, warnings)

	return &simRes, nil
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

func getIntensityConfig(intensity string, intensities []configs.Intensity) (*configs.Intensity, error) {
	for i := range intensities {
		if intensities[i].Type == intensity {
			return &intensities[i], nil
		}
	}
	return nil, fmt.Errorf("intensity %s not found", intensity)
}
