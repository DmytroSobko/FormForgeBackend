package dto

type SimulationConfigResponse struct {
	RestDayRecovery      float64 `json:"restDayRecovery"`
	MaxFatiguePenalty    float64 `json:"maxFatiguePenalty"`
	HighFatigueThreshold float64 `json:"highFatigueThreshold"`
}
