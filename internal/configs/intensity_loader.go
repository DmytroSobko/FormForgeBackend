package configs

import "fmt"

func LoadIntensities(path string) (*IntensityEnvelope, error) {
	var cfg IntensityEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, err
	}

	if cfg.Version == "" {
		return nil, fmt.Errorf("intensity config missing version")
	}

	for name, i := range cfg.Intensities {
		if err := i.Validate(name); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
