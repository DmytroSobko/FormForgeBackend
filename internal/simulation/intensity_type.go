package simulation

import (
	"encoding/json"
	"fmt"
)

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

func (i *IntensityType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t := IntensityType(s)
	if !t.IsValid() {
		return fmt.Errorf("invalid intensity type: %s", s)
	}

	*i = t
	return nil
}

func (i IntensityType) String() string {
	return string(i)
}
