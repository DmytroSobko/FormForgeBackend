package dto

type CreateAthleteRequest struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type AthleteResponse struct {
	ID         string  `json:"id"`
	Type       string  `json:"type"`
	Name       string  `json:"name"`
	Strength   float64 `json:"strength"`
	Endurance  float64 `json:"endurance"`
	Mobility   float64 `json:"mobility"`
	Fatigue    float64 `json:"fatigue"`
	MaxFatigue float64 `json:"maxFatigue"`
	Week       int     `json:"week"`
}
