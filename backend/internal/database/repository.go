package database

import (
	"context"
	"os"

	"backend/models"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
    db *pgx.Conn
}

func NewRepository() (*Repository, error) {
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        return nil, err
    }

    return &Repository{db: conn}, nil
}

func (repo *Repository) UpsertPingResults(results models.PingResults) error {
	query := `
		INSERT INTO ping_results (ip, ping_time, date)
		VALUES ($1, $2, $3)
		ON CONFLICT (ip) DO UPDATE
		SET ping_time = EXCLUDED.ping_time, date = EXCLUDED.date
	`

	for _, result := range results.Results {
		_, err := repo.db.Exec(context.Background(), query, result.IP, result.PingTime, result.Date)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *Repository) Close() {
    repo.db.Close(context.Background())
}