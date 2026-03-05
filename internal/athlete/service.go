package athlete

import (
	"context"
	"fmt"
	"log"

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
	log.Printf("CreateAthlete started: type=%s name=%s", athleteType, name)

	t, ok := s.athleteTypes[athleteType]
	if !ok {
		return nil, fmt.Errorf("Invalid athlete type: %s", athleteType)
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
		log.Println("NewAthlete failed")

		return nil, err
	}

	if err := s.repo.Save(ctx, athlete); err != nil {
		log.Println("Repo save failed")
		return nil, err
	}

	log.Println("CreateAthlete: athlete created successfully")

	return athlete, nil
}
