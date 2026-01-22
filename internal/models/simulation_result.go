package models

import "time"

type StatSnapshot struct {
	Strength  float64 `json:"strength"`
	Endurance float64 `json:"endurance"`
	Mobility  float64 `json:"mobility"`
	Fatigue   float64 `json:"fatigue"`
}

type SimulationResult struct {
	ID         string       `json:"id"`
	AthleteID  string       `json:"athleteId"`
	Week       int          `json:"week"`
	Before     StatSnapshot `json:"before"`
	After      StatSnapshot `json:"after"`
	Efficiency float64      `json:"efficiency"`
	Warnings   []string     `json:"warnings"`
	CreatedAt  time.Time    `json:"createdAt"`
}

// NewSimulationResult constructs a result object.
func NewSimulationResult(id string, athleteID string, week int, before StatSnapshot,
	after StatSnapshot, efficiency float64, warnings []string) SimulationResult {
	return SimulationResult{
		ID:         id,
		AthleteID:  athleteID,
		Week:       week,
		Before:     before,
		After:      after,
		Efficiency: efficiency,
		Warnings:   warnings,
		CreatedAt:  time.Now(),
	}
}
