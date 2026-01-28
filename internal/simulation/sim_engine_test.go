package simulation

import (
	"testing"

	"github.com/DmytroSobko/FormForgeBackend/internal/config"
	"github.com/DmytroSobko/FormForgeBackend/internal/models"
)

func testConfig() *config.SimulationConfig {
	return &config.SimulationConfig{
		DaysInWeek:           7,
		RestDayRecovery:      15.0,
		MaxFatiguePenalty:    0.7,
		HighFatigueThreshold: 0.6,
		IntensityMultipliers: map[string]float64{
			"low":    0.6,
			"medium": 1.0,
			"high":   1.4,
		},
	}
}

func baseAthlete() models.Athlete {
	return models.Athlete{
		ID:         "athlete-1",
		Strength:   50,
		Endurance:  50,
		Mobility:   50,
		Fatigue:    0,
		MaxFatigue: 100,
		Week:       1,
	}
}

func simpleTrainingPlan() models.TrainingPlan {
	days := make([]models.TrainingDay, 7)

	for i := 0; i < 7; i++ {
		days[i] = models.TrainingDay{
			DayIndex: i,
			Exercises: []models.Exercise{
				{
					BaseGain:    5,
					FatigueCost: 10,
					Intensity:   models.IntensityMedium,
					PrimaryStat: models.StatStrength,
				},
			},
		}
	}

	return models.TrainingPlan{Days: days}
}

func restOnlyPlan() models.TrainingPlan {
	days := make([]models.TrainingDay, 7)

	for i := 0; i < 7; i++ {
		days[i] = models.TrainingDay{
			DayIndex:  i,
			Exercises: []models.Exercise{},
		}
	}

	return models.TrainingPlan{Days: days}
}

// ------------------------------------------------
// Tests
// ------------------------------------------------

func TestSimulationIsDeterministic(t *testing.T) {
	cfg := testConfig()
	engine := NewEngine(cfg)

	a1 := baseAthlete()
	a2 := baseAthlete()
	plan := simpleTrainingPlan()

	r1 := engine.SimulateWeek(a1, plan)
	r2 := engine.SimulateWeek(a2, plan)

	if r1.After.Strength != r2.After.Strength {
		t.Errorf("strength mismatch: %v vs %v", r1.After.Strength, r2.After.Strength)
	}

	if r1.After.Fatigue != r2.After.Fatigue {
		t.Errorf("fatigue mismatch: %v vs %v", r1.After.Fatigue, r2.After.Fatigue)
	}
}

func TestFatigueIncreasesWithTraining(t *testing.T) {
	cfg := testConfig()
	engine := NewEngine(cfg)

	athlete := baseAthlete()
	plan := simpleTrainingPlan()

	result := engine.SimulateWeek(athlete, plan)

	if result.After.Fatigue <= 0 {
		t.Errorf("expected fatigue to increase, got %v", result.After.Fatigue)
	}
}

func TestRestDayReducesFatigue(t *testing.T) {
	cfg := testConfig()
	engine := NewEngine(cfg)

	athlete := baseAthlete()
	athlete.Fatigue = 50

	result := engine.SimulateWeek(athlete, restOnlyPlan())

	if result.After.Fatigue >= 50 {
		t.Errorf("expected fatigue to reduce, got %v", result.After.Fatigue)
	}
}

func TestHighFatigueReducesEfficiency(t *testing.T) {
	cfg := testConfig()
	engine := NewEngine(cfg)

	athlete := baseAthlete()
	athlete.Fatigue = 90

	result := engine.SimulateWeek(athlete, simpleTrainingPlan())

	if result.Efficiency >= 1.0 {
		t.Errorf("expected efficiency < 1.0, got %v", result.Efficiency)
	}
}

func TestHighFatigueWarningTriggered(t *testing.T) {
	cfg := testConfig()
	engine := NewEngine(cfg)

	athlete := baseAthlete()
	athlete.Fatigue = 90

	result := engine.SimulateWeek(athlete, simpleTrainingPlan())

	if len(result.Warnings) == 0 {
		t.Error("expected fatigue warning but got none")
	}
}
