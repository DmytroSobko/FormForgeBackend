package models

// TrainingPlan represents a weekly plan consisting of 7 days.
type TrainingPlan struct {
	ID        string        `json:"id"`
	AthleteID string        `json:"athleteId"`
	Days      []TrainingDay `json:"days"`
}

// TrainingDay represents a single day of training.
type TrainingDay struct {
	DayIndex  int        `json:"dayIndex"` // 0 = Monday, 6 = Sunday
	Exercises []Exercise `json:"exercises"`
}

// NewEmptyTrainingPlan creates a plan with 7 empty days.
func NewEmptyTrainingPlan(id string, athleteID string) TrainingPlan {
	days := make([]TrainingDay, DaysInWeek)

	for i := 0; i < DaysInWeek; i++ {
		days[i] = TrainingDay{
			DayIndex:  i,
			Exercises: []Exercise{},
		}
	}

	return TrainingPlan{
		ID:        id,
		AthleteID: athleteID,
		Days:      days,
	}
}
