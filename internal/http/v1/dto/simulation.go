package dto

type SimulateWeekRequest struct {
	AthleteID string              `json:"athleteId"`
	Plan      TrainingPlanRequest `json:"plan"`
}

type TrainingPlanRequest struct {
	Days []TrainingDayRequest `json:"days"`
}

type TrainingDayRequest struct {
	DayIndex  int                      `json:"dayIndex"`
	Exercises []PlannedExerciseRequest `json:"exercises"`
}

type PlannedExerciseRequest struct {
	ExerciseType string `json:"exerciseType"`
	Intensity    string `json:"intensity"`
}

type SimulationResultResponse struct {
	ID         string   `json:"id"`
	AthleteID  string   `json:"athleteId"`
	Week       int      `json:"week"`
	Efficiency float64  `json:"efficiency"`
	Warnings   []string `json:"warnings"`
}
