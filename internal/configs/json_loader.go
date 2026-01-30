package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadJSON(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file: %w", err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("parse json: %w", err)
	}

	return nil
}
