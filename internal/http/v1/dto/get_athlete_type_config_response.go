package dto

type StatBlock struct {
	Strength  float64 `json:"strength"`
	Endurance float64 `json:"endurance"`
	Mobility  float64 `json:"mobility"`
}

type AthleteTypeConfig struct {
	Type               string    `json:"type"`
	DisplayName        string    `json:"displayName"`
	Description        string    `json:"description"`
	BaseStats          StatBlock `json:"baseStats"`
	MaxFatigue         float64   `json:"maxFatigue"`
	RecoveryMultiplier float64   `json:"recoveryMultiplier"`
	FatigueSensitivity float64   `json:"fatigueSensitivity"`
}

type AthleteTypeConfigsResponse struct {
	AthleteTypes []AthleteTypeConfig `json:"athleteTypes"`
}
