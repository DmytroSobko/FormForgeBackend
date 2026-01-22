package models

type StatType string

const (
	StatStrength  StatType = "strength"
	StatEndurance StatType = "endurance"
	StatMobility  StatType = "mobility"
)

type IntensityType string

const (
	IntensityLow    IntensityType = "low"
	IntensityMedium IntensityType = "medium"
	IntensityHigh   IntensityType = "high"
)

type Exercise struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	PrimaryStat     StatType      `json:"primaryStat"`
	SecondaryStat   *StatType     `json:"secondaryStat,omitempty"`
	Intensity       IntensityType `json:"intensity"`
	BaseGain        float64       `json:"baseGain"`
	FatigueCost     float64       `json:"fatigueCost"`
	DurationMinutes int           `json:"durationMinutes"`
}
