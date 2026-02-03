package configs

import "fmt"

func LoadExercises(path string) (*ExercisesEnvelope, error) {
	var cfg ExercisesEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, err
	}

	if cfg.Version == "" {
		return nil, fmt.Errorf("exercise config missing version")
	}

	ids := map[string]bool{}
	for _, e := range cfg.Exercises {
		if ids[e.ID] {
			return nil, fmt.Errorf("duplicate exercise id: %s", e.ID)
		}
		ids[e.ID] = true

		if err := e.Validate(); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
