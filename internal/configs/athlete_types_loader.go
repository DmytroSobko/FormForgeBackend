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

	types := map[string]bool{}
	for _, a := range cfg.AthletesTypes {
		if types[a.Type] {
			return nil, fmt.Errorf("duplicate athlete type: %s", a.Type)
		}
		types[a.Type] = true

		if err := a.Validate(); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
