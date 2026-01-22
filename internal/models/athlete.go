package models

type Athlete struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Strength   float64 `json:"strength"`
	Endurance  float64 `json:"endurance"`
	Mobility   float64 `json:"mobility"`
	Fatigue    float64 `json:"fatigue"`
	MaxFatigue float64 `json:"maxFatigue"`
	Week       int     `json:"week"`
}

func NewAthlete(id string, name string, strength float64, endurance float64, mobility float64, maxFatigue float64) Athlete {
	return Athlete{
		ID:         id,
		Name:       name,
		Strength:   strength,
		Endurance:  endurance,
		Mobility:   mobility,
		Fatigue:    0,
		MaxFatigue: maxFatigue,
		Week:       1,
	}
}
