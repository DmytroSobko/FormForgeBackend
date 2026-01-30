package configs

import "fmt"

func LoadAthleteTypes(path string) (*AthleteTypesEnvelope, error) {
	var cfg AthleteTypesEnvelope

	if err := loadJSON(path, &cfg); err != nil {
		return nil, err
	}

	if cfg.Version == "" {
		return nil, fmt.Errorf("athlete config missing version")
	}

	ids := map[string]bool{}
	for _, a := range cfg.Athletes {
		if ids[a.ID] {
			return nil, fmt.Errorf("duplicate athlete id: %s", a.ID)
		}
		ids[a.ID] = true

		if err := a.Validate(); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
