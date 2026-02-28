package athlete

import (
	"context"

	"github.com/DmytroSobko/FormForgeBackend/internal/db"
)

type PostgresRepository struct {
	db *db.DB
}

func NewPostgresRepository(db *db.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(ctx context.Context, a *Athlete) error {

	query := `
		INSERT INTO athletes (
			id, type, name,
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
