package database

import (
	"context"
	"os"
	"log"
	"time"
	"fmt"

	"backend/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
    db *pgxpool.Pool
}

func NewRepository(maxRetries int, retryInterval time.Duration) (*Repository, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL не установлена")
	}

	var conn *pgxpool.Pool
	var err error

	for i := 1; i <= maxRetries; i++ {
		conn, err = pgxpool.Connect(context.Background(), databaseUrl)
		if err == nil {
			return &Repository{db: conn}, nil
		}

		log.Printf("Не удалось подключиться к PostgreSQL. Попытка %d из %d: %v", i, maxRetries, err)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("не удалось подключиться к PostgreSQL после %d попыток: %w", maxRetries, err)
}

func (repo *Repository) UpsertPingResults(results models.PingResults) error {

	tx, err := repo.db.BeginTx(context.Background(), pgx.TxOptions{})
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback(context.Background()) 
        }
    }()

    query := `
        INSERT INTO ping_results (ip, ping_time, date)
        VALUES ($1, $2, $3)
        ON CONFLICT (ip) DO UPDATE
        SET ping_time = EXCLUDED.ping_time, date = EXCLUDED.date
    `
    for _, result := range results.Results {
        _, err := tx.Exec(context.Background(), query, result.IP, result.PingTime, result.Date)
        if err != nil {
            return err
        }
    }

    return tx.Commit(context.Background())
}

func (repo *Repository) GetPingResults() (models.PingResults, error) {
    query := `SELECT ip, ping_time, date FROM ping_results`
	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return models.PingResults{}, err
	}
	defer rows.Close()

	var results models.PingResults
	for rows.Next() {
		var pr models.PingResult
		if err := rows.Scan(&pr.IP, &pr.PingTime, &pr.Date); err != nil {
			return models.PingResults{}, err
		}
		results.Results = append(results.Results, pr)
	}
	return results, nil
}

func (repo *Repository) Close() {
    repo.db.Close() 
}
