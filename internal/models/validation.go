package models

import (
	"errors"
	"fmt"
)

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

func (a Athlete) Validate() error {
	if a.ID == "" {
		return errors.New("athlete id cannot be empty")
	}

	if a.Name == "" {
		return errors.New("athlete name cannot be empty")
	}

	if err := validateStatRange("strength", a.Strength); err != nil {
		return err
	}
	if err := validateStatRange("endurance", a.Endurance); err != nil {
		return err
	}
	if err := validateStatRange("mobility", a.Mobility); err != nil {
		return err
	}

	if err := validateNonNegative("fatigue", a.Fatigue); err != nil {
		return err
	}

	if a.MaxFatigue <= 0 {
		return errors.New("maxFatigue must be greater than 0")
	}

	if a.Fatigue > a.MaxFatigue {
		return errors.New("fatigue cannot exceed maxFatigue")
	}

	if a.Week < 1 {
		return errors.New("week must be >= 1")
	}

	return nil
}

func (e Exercise) Validate() error {
	if e.ID == "" {
		return errors.New("exercise id cannot be empty")
	}

	if e.Name == "" {
		return errors.New("exercise name cannot be empty")
	}

	if e.BaseGain < 0 {
		return errors.New("baseGain must be >= 0")
	}

	if e.FatigueCost < 0 {
		return errors.New("fatigueCost must be >= 0")
	}

	if e.DurationMinutes <= 0 {
		return errors.New("durationMinutes must be > 0")
	}

	switch e.PrimaryStat {
	case StatStrength, StatEndurance, StatMobility:
		// ok
	default:
		return fmt.Errorf("invalid primaryStat: %s", e.PrimaryStat)
	}

	if e.SecondaryStat != nil {
		switch *e.SecondaryStat {
		case StatStrength, StatEndurance, StatMobility:
			// ok
		default:
			return fmt.Errorf("invalid secondaryStat: %s", *e.SecondaryStat)
		}
	}

	switch e.Intensity {
	case IntensityLow, IntensityMedium, IntensityHigh:
		// ok
	default:
		return fmt.Errorf("invalid intensity: %s", e.Intensity)
	}

	return nil
}

func (p TrainingPlan) Validate() error {
	if p.ID == "" {
		return errors.New("training plan id cannot be empty")
	}

	if p.AthleteID == "" {
		return errors.New("training plan athleteId cannot be empty")
	}

	if len(p.Days) != DaysInWeek {
		return errors.New("training plan must contain exactly 7 days")
	}

	for _, day := range p.Days {
		if day.DayIndex < 0 || day.DayIndex > DaysInWeek-1 {
			return fmt.Errorf("invalid day index: %d", day.DayIndex)
		}

		for _, exercise := range day.Exercises {
			if err := exercise.Validate(); err != nil {
				return fmt.Errorf("day %d: %w", day.DayIndex, err)
			}
		}
	}

	return nil
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
