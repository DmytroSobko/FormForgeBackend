package athlete

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, a *Athlete) error
}

type Service struct {
	repo         Repository
	athleteTypes map[AthleteType]AthleteTypeConfig
}

func NewService(repo Repository, types []AthleteTypeConfig) *Service {
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

	t, ok := s.athleteTypes[athleteType]
	if !ok {
		return nil, fmt.Errorf("invalid athlete type: %s", athleteType)
	}

	id := uuid.NewString()

	athlete, err := NewAthlete(
		id,
		t.Type,
		name,
		t.BaseStats.Strength.Value(),
		t.BaseStats.Endurance.Value(),
		t.BaseStats.Mobility.Value(),
		t.MaxFatigue,
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, athlete); err != nil {
		return nil, err
	}

	return athlete, nil
}
