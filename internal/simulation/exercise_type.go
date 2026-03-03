package simulation

import (
	"encoding/json"
	"fmt"
)

type ExerciseType string

const (
	BenchPress    ExerciseType = "bench_press"
	Squat         ExerciseType = "squat"
	Deadlift      ExerciseType = "deadlift"
	OverheadPress ExerciseType = "overhead_press"
	Running       ExerciseType = "running"
	Cycling       ExerciseType = "cycling"
	Rowing        ExerciseType = "rowing"
	Stretching    ExerciseType = "stretching"
	YogaFlow      ExerciseType = "yoga_flow"
	CoreStability ExerciseType = "core_stability"
)

func (i ExerciseType) IsValid() bool {
	switch i {
	case BenchPress, Squat, Deadlift, OverheadPress, Running, Cycling, Rowing, Stretching, YogaFlow, CoreStability:
		return true
	default:
		return false
	}
}

func (i *ExerciseType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t := ExerciseType(s)
	if !t.IsValid() {
		return fmt.Errorf("invalid exercise type: %s", s)
	}

	*i = t
	return nil
}

func (i ExerciseType) String() string {
	return string(i)
}
