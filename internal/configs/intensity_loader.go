package configs

import "fmt"

func LoadIntensities(path string) (*IntensitiesEnvelope, error) {
	var cfg IntensitiesEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, err
	}

	if cfg.Version == "" {
		return nil, fmt.Errorf("intensity config missing version")
	}

	types := map[string]bool{}
	for _, i := range cfg.Intensities {
		if types[i.Type] {
			return nil, fmt.Errorf("duplicate intensity type: %s", i.Type)
		}
		types[i.Type] = true

		if err := i.Validate(); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
