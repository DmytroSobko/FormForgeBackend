package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "simulation.json")

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	return path
}

func TestLoadSimulationConfig_Success(t *testing.T) {
	json := `
{
  "version": "1.0.0",
  "simulation": {
    "daysInWeek": 7,
    "restDayRecovery": 15.0,
    "maxFatiguePenalty": 0.7,
    "highFatigueThreshold": 0.6,
    "intensityMultipliers": {
      "low": 0.6,
      "medium": 1.0,
      "high": 1.4
    }
  }
}
`

	path := writeTempConfig(t, json)

	cfg, err := LoadSimulationConfig(path)
	if err != nil {
		t.Fatalf("expected config to load, got error: %v", err)
	}

	if cfg.Version != "1.0.0" {
		t.Errorf("unexpected version: %s", cfg.Version)
	}

	if cfg.Simulation.DaysInWeek != 7 {
		t.Errorf("unexpected daysInWeek: %d", cfg.Simulation.DaysInWeek)
	}
}

func TestLoadSimulationConfig_FailsOnInvalidDays(t *testing.T) {
	json := `
{
  "version": "1.0.0",
  "simulation": {
    "daysInWeek": 0,
    "restDayRecovery": 15.0,
    "maxFatiguePenalty": 0.7,
    "highFatigueThreshold": 0.6,
    "intensityMultipliers": {
      "low": 0.6,
      "medium": 1.0,
      "high": 1.4
    }
  }
}
`

	path := writeTempConfig(t, json)

	if _, err := LoadSimulationConfig(path); err == nil {
		t.Error("expected error for invalid daysInWeek, got nil")
	}
}

func TestLoadSimulationConfig_FailsOnMissingIntensity(t *testing.T) {
	json := `
{
  "version": "1.0.0",
  "simulation": {
    "daysInWeek": 7,
    "restDayRecovery": 15.0,
    "maxFatiguePenalty": 0.7,
    "highFatigueThreshold": 0.6,
    "intensityMultipliers": {
      "low": 0.6,
      "medium": 1.0
    }
  }
}
`

	path := writeTempConfig(t, json)

	if _, err := LoadSimulationConfig(path); err == nil {
		t.Error("expected error for missing intensity multiplier")
	}
}

func TestLoadSimulationConfig_FailsOnInvalidThresholds(t *testing.T) {
	json := `
{
  "version": "1.0.0",
  "simulation": {
    "daysInWeek": 7,
    "restDayRecovery": 15.0,
    "maxFatiguePenalty": 0.5,
    "highFatigueThreshold": 0.9,
    "intensityMultipliers": {
      "low": 0.6,
      "medium": 1.0,
      "high": 1.4
    }
  }
}
`

	path := writeTempConfig(t, json)

	if _, err := LoadSimulationConfig(path); err == nil {
		t.Error("expected error for highFatigueThreshold > maxFatiguePenalty")
	}
}
