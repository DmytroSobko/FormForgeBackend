package simulation

import (
	"testing"

	configs "github.com/DmytroSobko/FormForgeBackend/internal/configs"
	models "github.com/DmytroSobko/FormForgeBackend/internal/models"
)

// --------------------
// Test helpers
// --------------------

func testSimulationConfig() *configs.SimulationConfig {
	return &configs.SimulationConfig{
		RestDayRecovery:      15.0,
		MaxFatiguePenalty:    0.7,
		HighFatigueThreshold: 0.6,
	}
}

func testIntensities() map[models.IntensityType]configs.Intensity {
	return map[models.IntensityType]configs.Intensity{
		models.IntensityLow: {
			Multiplier:        0.6,
			FatigueMultiplier: 0.6,
		},
		models.IntensityMedium: {
			Multiplier:        1.0,
			FatigueMultiplier: 1.0,
		},
		models.IntensityHigh: {
			Multiplier:        1.4,
			FatigueMultiplier: 1.4,
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

func baseExercise() models.Exercise {
	return models.Exercise{
		ID:          "squat",
		PrimaryStat: models.StatStrength,
		BaseGain:    5,
		FatigueCost: 10,
		AllowedIntensities: []string{
			"low",
			"medium",
			"high",
		},
	}
}

func simpleTrainingPlan() models.TrainingPlan {
	days := make([]models.TrainingDay, models.DaysInWeek)

	ex := baseExercise()

	for i := 0; i < models.DaysInWeek; i++ {
		days[i] = models.TrainingDay{
			DayIndex: i,
			Exercises: []models.PlannedExercise{
				{
					Exercise:  ex,
					Intensity: models.IntensityMedium,
				},
			},
		}
	}

	return models.TrainingPlan{
		ID:        "plan-1",
		AthleteID: "athlete-1",
		Days:      days,
	}
}

func restOnlyPlan() models.TrainingPlan {
	days := make([]models.TrainingDay, models.DaysInWeek)

	for i := 0; i < models.DaysInWeek; i++ {
		days[i] = models.TrainingDay{
			DayIndex:  i,
			Exercises: []models.PlannedExercise{},
		}
	}

	return models.TrainingPlan{
		ID:        "rest-plan",
		AthleteID: "athlete-1",
		Days:      days,
	}
}

// --------------------
// Tests
// --------------------

func TestSimulationIsDeterministic(t *testing.T) {
	engine := NewEngine(
		testSimulationConfig(),
		testIntensities(),
	)

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
	engine := NewEngine(
		testSimulationConfig(),
		testIntensities(),
	)

	athlete := baseAthlete()
	plan := simpleTrainingPlan()

	result := engine.SimulateWeek(athlete, plan)

	if result.After.Fatigue <= 0 {
		t.Errorf("expected fatigue to increase, got %v", result.After.Fatigue)
	}
}

func TestRestDayReducesFatigue(t *testing.T) {
	engine := NewEngine(
		testSimulationConfig(),
		testIntensities(),
	)

	athlete := baseAthlete()
	athlete.Fatigue = 50

	result := engine.SimulateWeek(athlete, restOnlyPlan())

	if result.After.Fatigue >= 50 {
		t.Errorf("expected fatigue to reduce, got %v", result.After.Fatigue)
	}
}

func TestHighFatigueReducesEfficiency(t *testing.T) {
	engine := NewEngine(
		testSimulationConfig(),
		testIntensities(),
	)

	athlete := baseAthlete()
	athlete.Fatigue = 90

	result := engine.SimulateWeek(athlete, simpleTrainingPlan())

	if result.Efficiency >= 1.0 {
		t.Errorf("expected efficiency < 1.0, got %v", result.Efficiency)
	}
}

func TestHighFatigueWarningTriggered(t *testing.T) {
	engine := NewEngine(
		testSimulationConfig(),
		testIntensities(),
	)

	athlete := baseAthlete()
	athlete.Fatigue = 90

	result := engine.SimulateWeek(athlete, simpleTrainingPlan())

	if len(result.Warnings) == 0 {
		t.Error("expected fatigue warning but got none")
	}
}
