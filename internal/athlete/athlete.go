package athlete

import (
	"errors"
	"fmt"
)

type Athlete struct {
	ID         string
	Type       AthleteType
	Name       string
	Strength   Stat
	Endurance  Stat
	Mobility   Stat
	Fatigue    float64
	MaxFatigue float64
	Week       int
}

func NewAthlete(
	id string,
	athleteType AthleteType,
	name string,
	strength float64,
	endurance float64,
	mobility float64,
	maxFatigue float64,
) (*Athlete, error) {

	if id == "" {
		return nil, errors.New("athlete id cannot be empty")
	}

	if name == "" {
		return nil, errors.New("athlete name cannot be empty")
	}

	if !athleteType.IsValid() {
		return nil, fmt.Errorf("invalid athlete type: %s", athleteType)
	}

	str, err := NewStat(strength)
	if err != nil {
		return nil, err
	}

	end, err := NewStat(endurance)
	if err != nil {
		return nil, err
	}

	mob, err := NewStat(mobility)
	if err != nil {
		return nil, err
	}

	if maxFatigue <= 0 {
		return nil, errors.New("maxFatigue must be greater than 0")
	}

	return &Athlete{
		ID:         id,
		Type:       athleteType,
		Name:       name,
		Strength:   str,
		Endurance:  end,
		Mobility:   mob,
		Fatigue:    0,
		MaxFatigue: maxFatigue,
		Week:       1,
	}, nil
}

func (a *Athlete) GetID() string {
	return a.ID
}

func (a *Athlete) GetType() AthleteType {
	return a.Type
}

func (a *Athlete) GetName() string {
	return a.Name
}

func (a *Athlete) GetWeek() int {
	return a.Week
}

func (a *Athlete) GetStrength() Stat {
	return a.Strength
}

func (a *Athlete) GetEndurance() Stat {
	return a.Endurance
}

func (a *Athlete) GetMobility() Stat {
	return a.Mobility
}

func (a *Athlete) GetFatigue() float64 {
	return a.Fatigue
}

func (a *Athlete) GetMaxFatigue() float64 {
	return a.MaxFatigue
}

func (a *Athlete) FatigueRatio() float64 {
	if a.MaxFatigue == 0 {
		return 0
	}
	return a.Fatigue / a.MaxFatigue
}

func (a *Athlete) ReduceFatigue(amount float64) {
	if amount < 0 {
		return
	}
	a.Fatigue -= amount
	if a.Fatigue < 0 {
		a.Fatigue = 0
	}
}

func (a *Athlete) AddFatigue(amount float64) error {
	if amount < 0 {
		return errors.New("fatigue increase must be non-negative")
	}

	a.Fatigue += amount

	if a.Fatigue > a.MaxFatigue {
		a.Fatigue = a.MaxFatigue
	}

	return nil
}

func (a *Athlete) ApplyStat(statType StatType, value float64) error {
	if value < 0 {
		return errors.New("stat increase must be non-negative")
	}

	current := float64(0)

	switch statType {
	case StatStrength:
		current = a.Strength.Value()
	case StatEndurance:
		current = a.Endurance.Value()
	case StatMobility:
		current = a.Mobility.Value()
	default:
		return errors.New("invalid stat type")
	}

	newValue := current + value

	stat, err := NewStat(newValue)
	if err != nil {
		return err
	}

	switch statType {
	case StatStrength:
		a.Strength = stat
	case StatEndurance:
		a.Endurance = stat
	case StatMobility:
		a.Mobility = stat
	}

	return nil
}
