package athlete

import (
	"encoding/json"
	"fmt"
)

type AthleteType string

const (
	Balanced         AthleteType = "balanced"
	EnduranceFocused AthleteType = "endurance_focused"
	StrengthFocused  AthleteType = "strength_focused"
)

func (a AthleteType) IsValid() bool {
	switch a {
	case Balanced, EnduranceFocused, StrengthFocused:
		return true
	default:
		return false
	}
}

func (a *AthleteType) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	t := AthleteType(str)
	if !t.IsValid() {
		return fmt.Errorf("invalid athlete type: %s", str)
	}

	*a = t
	return nil
}

func (a AthleteType) String() string {
	return string(a)
}
