package models

import "fmt"

type StatType string

const (
	StatStrength  StatType = "strength"
	StatEndurance StatType = "endurance"
	StatMobility  StatType = "mobility"
)

func (s StatType) IsValid() bool {
	switch s {
	case StatStrength, StatEndurance, StatMobility:
		return true
	default:
		return false
	}
}

func (s *StatType) UnmarshalText(text []byte) error {
	value := StatType(text)

	if !value.IsValid() {
		return fmt.Errorf("invalid stat type: %s", text)
	}

	*s = value
	return nil
}
