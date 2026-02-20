package models

import (
	"errors"
	"fmt"
)

// TrainingPlan represents a weekly plan consisting of 7 days.
type TrainingPlan struct {
	ID        string        `json:"id"`
	AthleteID string        `json:"athleteId"`
	Days      []TrainingDay `json:"days"`
}

// TrainingDay represents a single day of training.
type TrainingDay struct {
	DayIndex  int               `json:"dayIndex"` // 0 = Monday, 6 = Sunday
	Exercises []PlannedExercise `json:"exercises"`
}

// PlannedExercise represents an exercise scheduled with a chosen intensity.
type PlannedExercise struct {
	Exercise  Exercise `json:"exercise"`
	Intensity string   `json:"intensity"`
}

// NewEmptyTrainingPlan creates a plan with 7 empty days.
func NewEmptyTrainingPlan(id string, athleteID string) TrainingPlan {
	days := make([]TrainingDay, DaysInWeek)

	for i := 0; i < DaysInWeek; i++ {
		days[i] = TrainingDay{
			DayIndex:  i,
			Exercises: []PlannedExercise{},
		}
	}

	return TrainingPlan{
		ID:        id,
		AthleteID: athleteID,
		Days:      days,
	}
}

// Validate ensures the training plan is structurally and logically correct.
func (p TrainingPlan) Validate() error {
	if p.ID == "" {
		return errors.New("training plan id cannot be empty")
	}

	if p.AthleteID == "" {
		return errors.New("training plan athleteId cannot be empty")
	}

	if len(p.Days) != DaysInWeek {
		return fmt.Errorf("training plan must contain exactly %d days", DaysInWeek)
	}

	for _, day := range p.Days {
		if day.DayIndex < 0 || day.DayIndex > DaysInWeek-1 {
			return fmt.Errorf("invalid day index: %d", day.DayIndex)
		}

		for _, planned := range day.Exercises {

			// Validate exercise definition
			if err := planned.Exercise.Validate(); err != nil {
				return fmt.Errorf(
					"day %d exercise %s: %w",
					day.DayIndex,
					planned.Exercise.Type,
					err,
				)
			}
		}
	}

	return nil
}
