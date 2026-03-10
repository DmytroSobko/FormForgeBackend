package athlete

import (
	"context"

	"github.com/DmytroSobko/FormForgeBackend/internal/db"
	"github.com/DmytroSobko/FormForgeBackend/internal/logging"
)

type PostgresRepository struct {
	db *db.DB
}

func NewPostgresRepository(db *db.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(ctx context.Context, a *Athlete) error {

	logger := logging.FromContext(ctx)

	logger.Info("saving athlete", "id", a.GetID())

	query := `
		INSERT INTO athletes (
			id, athlete_type, name,
			strength, endurance, mobility,
			fatigue, max_fatigue, week
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	return r.db.Exec(
		ctx,
		query,
		a.GetID(),
		a.GetType().String(),
		a.GetName(),
		a.GetStrength().Value(),
		a.GetEndurance().Value(),
		a.GetMobility().Value(),
		a.GetFatigue(),
		a.GetMaxFatigue(),
		a.GetWeek(),
	)
}

func (r *PostgresRepository) GetAthletes(ctx context.Context, limit, offset int) ([]*Athlete, error) {
	logger := logging.FromContext(ctx)

	logger.Info("fetching athletes", "limit", limit, "offset", offset)

	query := `
		SELECT 
			id,
			athlete_type,
			name,
			strength,
			endurance,
			mobility,
			fatigue,
			max_fatigue,
			week
		FROM athletes
		ORDER BY name
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		logger.Error("failed to query athletes", "error", err)
		return nil, err
	}
	defer rows.Close()

	var athletes []*Athlete

	for rows.Next() {
		var (
			id         string
			typeStr    string
			name       string
			strength   float64
			endurance  float64
			mobility   float64
			fatigue    float64
			maxFatigue float64
			week       int
		)

		if err := rows.Scan(
			&id,
			&typeStr,
			&name,
			&strength,
			&endurance,
			&mobility,
			&fatigue,
			&maxFatigue,
			&week,
		); err != nil {
			return nil, err
		}

		a, err := NewAthlete(
			id,
			AthleteType(typeStr),
			name,
			strength,
			endurance,
			mobility,
			maxFatigue,
		)
		if err != nil {
			return nil, err
		}

		a.Fatigue = fatigue
		a.Week = week

		athletes = append(athletes, a)
	}

	return athletes, rows.Err()
}
