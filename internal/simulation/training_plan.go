package simulation

import (
	"errors"
	"fmt"
)

const DaysInWeek = 7

// ------------------------------------------------------
// Domain Types
// ------------------------------------------------------

type TrainingPlan struct {
	ID        string
	AthleteID string
	Days      []TrainingDay
}

type TrainingDay struct {
	DayIndex  int
	Exercises []PlannedExercise
}

type PlannedExercise struct {
	Exercise  Exercise
	Intensity IntensityType
}

// ------------------------------------------------------
// Constructors
// ------------------------------------------------------

func NewEmptyTrainingPlan(id string, athleteID string) (*TrainingPlan, error) {

	if id == "" {
		return nil, errors.New("training plan id cannot be empty")
	}

	if athleteID == "" {
		return nil, errors.New("training plan athleteId cannot be empty")
	}

	days := make([]TrainingDay, DaysInWeek)

	for i := 0; i < DaysInWeek; i++ {
		days[i] = TrainingDay{
			DayIndex:  i,
			Exercises: []PlannedExercise{},
		}
	}

	return &TrainingPlan{
		ID:        id,
		AthleteID: athleteID,
		Days:      days,
	}, nil
}

// ------------------------------------------------------
// Getters
// ------------------------------------------------------

func (p *TrainingPlan) GetID() string {
	return p.ID
}

func (p *TrainingPlan) GetAthleteID() string {
	return p.AthleteID
}

func (p *TrainingPlan) GetDays() []TrainingDay {
	return p.Days
}

func (d *TrainingDay) GetDayIndex() int {
	return d.DayIndex
}

func (d *TrainingDay) GetExercises() []PlannedExercise {
	return d.Exercises
}

// ------------------------------------------------------
// Behavior
// ------------------------------------------------------

func (p *TrainingPlan) AddExercise(dayIndex int, exercise Exercise, intensity IntensityType) error {

	if dayIndex < 0 || dayIndex >= DaysInWeek {
		return fmt.Errorf("invalid day index: %d", dayIndex)
	}

	if !intensity.IsValid() {
		return fmt.Errorf("invalid intensity: %s", intensity)
	}

	p.Days[dayIndex].Exercises = append(
		p.Days[dayIndex].Exercises,
		PlannedExercise{
			Exercise:  exercise,
			Intensity: intensity,
		},
	)

	return nil
}
