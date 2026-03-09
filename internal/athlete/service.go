package athlete

import (
	"context"
	"fmt"

	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, a *Athlete) error
}

type AthleteService interface {
	CreateAthlete(ctx context.Context, athleteType AthleteType, name string) (*Athlete, error)
}

type Service struct {
	repo         Repository
	athleteTypes map[AthleteType]AthleteTypeConfig
}

func NewService(repo Repository, types []AthleteTypeConfig) AthleteService {
	m := make(map[AthleteType]AthleteTypeConfig)
	for _, t := range types {
		m[t.Type] = t
	}

	return &Service{
		repo:         repo,
		athleteTypes: m,
	}
}

func (s *Service) CreateAthlete(ctx context.Context, athleteType AthleteType, name string) (*Athlete, error) {
	logger := logging.FromContext(ctx)

	logger.Info(
		"create athlete started",
		"type", athleteType,
		"name", name,
	)

	cfg, ok := s.athleteTypes[athleteType]
	if !ok {
		err := fmt.Errorf("unknown athlete type: %s", athleteType)

		logger.Warn(
			"invalid athlete type requested",
			"type", athleteType,
		)

		return nil, err
	}

	id := uuid.NewString()

	athlete, err := NewAthlete(
		id,
		cfg.Type,
		name,
		cfg.BaseStats.Strength.Value(),
		cfg.BaseStats.Endurance.Value(),
		cfg.BaseStats.Mobility.Value(),
		cfg.MaxFatigue,
	)

	if err != nil {
		logger.Error(
			"failed to create athlete entity",
			"error", err,
			"type", athleteType,
			"name", name,
		)
		return nil, err
	}

	if err := s.repo.Save(ctx, athlete); err != nil {
		logger.Error(
			"failed to persist athlete",
			"error", err,
			"id", athlete.ID,
		)
		return nil, err
	}

	logger.Info(
		"athlete created",
		"id", athlete.ID,
		"type", athlete.Type,
		"name", athlete.Name,
	)

	return athlete, nil
}
