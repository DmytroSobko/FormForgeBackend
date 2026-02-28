package simulation

import (
	"errors"
	"fmt"
	"time"

	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
)

type StatSnapshot struct {
	Strength  athlete.Stat `json:"strength"`
	Endurance athlete.Stat `json:"endurance"`
	Mobility  athlete.Stat `json:"mobility"`
	Fatigue   float64      `json:"fatigue"`
}

type SimulationResult struct {
	ID         string
	AthleteID  string
	Week       int
	Before     StatSnapshot
	After      StatSnapshot
	Efficiency float64
	Warnings   []string
	CreatedAt  time.Time
}

// ------------------------------------------------------
// Base constructor (low-level)
// ------------------------------------------------------

func NewSimulationResult(
	id string,
	athleteID string,
	week int,
	before StatSnapshot,
	after StatSnapshot,
	efficiency float64,
	warnings []string,
) (*SimulationResult, error) {

	if id == "" {
		return nil, errors.New("simulation result id cannot be empty")
	}

	if athleteID == "" {
		return nil, errors.New("simulation result athleteId cannot be empty")
	}

	if week < 1 {
		return nil, errors.New("week must be >= 1")
	}

	if efficiency < 0 {
		return nil, errors.New("efficiency must be >= 0")
	}

	return &SimulationResult{
		ID:         id,
		AthleteID:  athleteID,
		Week:       week,
		Before:     before,
		After:      after,
		Efficiency: efficiency,
		Warnings:   warnings,
		CreatedAt:  time.Now(),
	}, nil
}

// ------------------------------------------------------
// High-level constructor from Athlete (preferred)
// ------------------------------------------------------

func NewSimulationResultFromAthlete(
	id string,
	current *athlete.Athlete,
	before athlete.Athlete,
	efficiency float64,
	warnings []string,
) (*SimulationResult, error) {

	if id == "" {
		return nil, errors.New("simulation result id cannot be empty")
	}

	if efficiency < 0 {
		return nil, fmt.Errorf("efficiency must be >= 0")
	}

	beforeSnapshot := newSnapshot(&before)
	afterSnapshot := newSnapshot(current)

	return &SimulationResult{
		ID:         id,
		AthleteID:  current.GetID(),
		Week:       current.GetWeek(),
		Before:     beforeSnapshot,
		After:      afterSnapshot,
		Efficiency: efficiency,
		Warnings:   warnings,
		CreatedAt:  time.Now(),
	}, nil
}

// ------------------------------------------------------
// Internal helper (private)
// ------------------------------------------------------

func newSnapshot(a *athlete.Athlete) StatSnapshot {
	return StatSnapshot{
		Strength:  a.GetStrength(),
		Endurance: a.GetEndurance(),
		Mobility:  a.GetMobility(),
		Fatigue:   a.GetFatigue(),
	}
}
