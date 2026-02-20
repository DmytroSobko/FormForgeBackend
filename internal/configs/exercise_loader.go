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

	types := map[string]bool{}
	for _, e := range cfg.Exercises {
		if types[e.Type] {
			return nil, fmt.Errorf("duplicate exercise id: %s", e.Type)
		}
		types[e.Type] = true

		if err := e.Validate(); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
