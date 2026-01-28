package config

import (
	"encoding/json"
	"os"
)

func LoadSimulationConfig(path string) (*SimulationConfigEnvelope, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var envelope SimulationConfigEnvelope
	if err := json.Unmarshal(data, &envelope); err != nil {
		return nil, err
	}

	if err := envelope.Simulation.Validate(); err != nil {
		return nil, err
	}

	return &envelope, nil
}
