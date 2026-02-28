package simulation

import (
	"fmt"
	"math"

	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
)

type Engine struct {
	cfg         Config
	intensities []Intensity
}

func NewEngine(
	cfg Config,
	intensities []Intensity,
) *Engine {
	return &Engine{
		cfg:         cfg,
		intensities: intensities,
	}
}

func (e *Engine) SimulateWeek(
	a athlete.Athlete,
	plan TrainingPlan,
) (*SimulationResult, error) {

	before := snapshot(&a)

	totalPotential := 0.0
	totalActual := 0.0
	var warnings []string

	for _, day := range plan.Days {

		// Rest day
		if len(day.Exercises) == 0 {
			a.ReduceFatigue(e.cfg.RestDayRecovery)
			continue
		}

		for _, planned := range day.Exercises {

			ex := planned.Exercise

			intensityCfg, err := e.getIntensity(planned.Intensity)
			if err != nil {
				return nil, err
			}

			fatigueRatio := a.FatigueRatio()
			penalty := math.Min(fatigueRatio, e.cfg.MaxFatiguePenalty)

			rawGain := ex.BaseGain * intensityCfg.Multiplier
			finalGain := rawGain * (1 - penalty)

			a.ApplyStat(ex.PrimaryStat, finalGain)

			a.AddFatigue(ex.FatigueCost * intensityCfg.FatigueFactor)

			totalPotential += rawGain
			totalActual += finalGain

			if penalty > e.cfg.HighFatigueThreshold {
				warnings = append(warnings, "High fatigue reduced gains")
			}
		}
	}

	after := snapshot(&a)

	efficiency := 1.0
	if totalPotential > 0 {
		efficiency = totalActual / totalPotential
	}

	return NewSimulationResult(
		"", // ID assigned later
		a.GetID(),
		a.GetWeek(),
		before,
		after,
		efficiency,
		warnings,
	)
}

func (e *Engine) getIntensity(t IntensityType) (Intensity, error) {
	for _, i := range e.intensities {
		if i.Type == t {
			return i, nil
		}
	}
	return Intensity{}, fmt.Errorf("intensity %s not found", t)
}

func snapshot(a *athlete.Athlete) StatSnapshot {
	return StatSnapshot{
		Strength:  a.GetStrength(),
		Endurance: a.GetEndurance(),
		Mobility:  a.GetMobility(),
		Fatigue:   a.GetFatigue(),
	}
}
