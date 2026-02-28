package mappers

import (
	"github.com/DmytroSobko/FormForgeBackend/internal/http/v1/dto"
	"github.com/DmytroSobko/FormForgeBackend/internal/simulation"
)

func ToSimulationResultResponse(r *simulation.SimulationResult) dto.SimulationResultResponse {
	return dto.SimulationResultResponse{
		ID:         r.ID,
		AthleteID:  r.AthleteID,
		Week:       r.Week,
		Efficiency: r.Efficiency,
		Warnings:   r.Warnings,
	}
}
