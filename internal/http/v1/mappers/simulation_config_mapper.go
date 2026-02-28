package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func ToSimulationConfigResponse(cfg simulation.Config) dto.SimulationConfigResponse {
	return dto.SimulationConfigResponse{
		RestDayRecovery:      cfg.RestDayRecovery,
		MaxFatiguePenalty:    cfg.MaxFatiguePenalty,
		HighFatigueThreshold: cfg.HighFatigueThreshold,
	}
}
