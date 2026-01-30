package configs

import "fmt"

func LoadSimulationConfig(path string) (*SimulationConfigEnvelope, error) {
	var envelope SimulationConfigEnvelope

	if err := loadJSON(path, &envelope); err != nil {
		return nil, err
	}

	if envelope.Version == "" {
		return nil, fmt.Errorf("simulation config missing version")
	}

	if err := envelope.Simulation.Validate(); err != nil {
		return nil, err
	}

	return &envelope, nil
}
