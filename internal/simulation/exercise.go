package simulation

import (
	"fmt"

	"github.com/DmytroSobko/FormForgeBackend/internal/athlete"
)

type Exercise struct {
	Type                ExerciseType
	DisplayName         string
	Description         string
	PrimaryStat         athlete.StatType
	SecondaryStat       *athlete.StatType
	SecondaryStatWeight float64
	BaseGain            float64
	FatigueCost         float64
	DurationMinutes     int
}

func NewExercise(
	exType string,
	displayName string,
	description string,
	primaryStat string,
	secondaryStat *string,
	secondaryWeight float64,
	baseGain float64,
	fatigueCost float64,
	durationMinutes int,
) (Exercise, error) {

	t := ExerciseType(exType)

	if !t.IsValid() {
		return Exercise{}, fmt.Errorf("invalid exercise type: %s", exType)
	}

	ps := athlete.StatType(primaryStat)
	if !ps.IsValid() {
		return Exercise{}, fmt.Errorf("invalid primaryStat: %s", primaryStat)
	}

	var ss *athlete.StatType
	if secondaryStat != nil {
		s := athlete.StatType(*secondaryStat)
		if !s.IsValid() {
			return Exercise{}, fmt.Errorf("invalid secondaryStat: %s", *secondaryStat)
		}
		ss = &s
	}

	if secondaryWeight < 0 || secondaryWeight > 1 {
		return Exercise{}, fmt.Errorf("secondaryStatWeight must be [0,1]")
	}

	if baseGain <= 0 {
		return Exercise{}, fmt.Errorf("baseGain must be > 0")
	}

	if fatigueCost < 0 {
		return Exercise{}, fmt.Errorf("fatigueCost must be >= 0")
	}

	if durationMinutes <= 0 {
		return Exercise{}, fmt.Errorf("durationMinutes must be > 0")
	}

	return Exercise{
		Type:                t,
		DisplayName:         displayName,
		Description:         description,
		PrimaryStat:         ps,
		SecondaryStat:       ss,
		SecondaryStatWeight: secondaryWeight,
		BaseGain:            baseGain,
		FatigueCost:         fatigueCost,
		DurationMinutes:     durationMinutes,
	}, nil
}
