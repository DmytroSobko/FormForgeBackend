package athlete

import (
	"encoding/json"
	"fmt"
)

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

func (s *StatType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	t := StatType(str)
	if !t.IsValid() {
		return fmt.Errorf("invalid stat type: %s", str)
	}

	*s = t
	return nil
}

func (s StatType) String() string {
	return string(s)
}
