package dto

type ExerciseConfig struct {
	Type                string  `json:"type"`
	DisplayName         string  `json:"displayName"`
	Description         string  `json:"description"`
	PrimaryStat         string  `json:"primaryStat"`
	SecondaryStat       *string `json:"secondaryStat,omitempty"`
	SecondaryStatWeight float64 `json:"secondaryStatWeight"`
	BaseGain            float64 `json:"baseGain"`
	FatigueCost         float64 `json:"fatigueCost"`
	DurationMinutes     int     `json:"durationMinutes"`
}

type ExerciseConfigsResponse struct {
	Exercises []ExerciseConfig `json:"exercises"`
}
