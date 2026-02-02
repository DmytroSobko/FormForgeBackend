package models

import "fmt"

type IntensityType string

const (
	IntensityLow    IntensityType = "low"
	IntensityMedium IntensityType = "medium"
	IntensityHigh   IntensityType = "high"
)

func (i IntensityType) IsValid() bool {
	switch i {
	case IntensityLow, IntensityMedium, IntensityHigh:
		return true
	default:
		return false
	}
}

func (i *IntensityType) UnmarshalText(text []byte) error {
	value := IntensityType(text)

	if !value.IsValid() {
		return fmt.Errorf("invalid intensity type: %s", text)
	}

	*i = value
	return nil
}
