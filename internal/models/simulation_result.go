package models

import (
	"errors"
	"fmt"
	"time"
)

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

func (r SimulationResult) Validate() error {
	if r.ID == "" {
		return errors.New("simulation result id cannot be empty")
	}

	if r.AthleteID == "" {
		return errors.New("simulation result athleteId cannot be empty")
	}

	if r.Week < 1 {
		return errors.New("simulation result week must be >= 1")
	}

	if err := validateStatRange("before.strength", r.Before.Strength); err != nil {
		return err
	}
	if err := validateStatRange("after.strength", r.After.Strength); err != nil {
		return err
	}

	if r.Efficiency < 0 {
		return errors.New("efficiency must be >= 0")
	}

	return nil
}

func validateStatRange(name string, value float64) error {
	if value < 0 || value > 100 {
		return fmt.Errorf("%s must be between 0 and 100 (got %.2f)", name, value)
	}
	return nil
}

func validateNonNegative(name string, value float64) error {
	if value < 0 {
		return fmt.Errorf("%s must be non-negative (got %.2f)", name, value)
	}
	return nil
}
