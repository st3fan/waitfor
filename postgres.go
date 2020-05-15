package waitfor

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresPool waits forever until Postgres is available. Keeps trying and never returns an error.
func PostgresPool(connString string) (*pgxpool.Pool, error) {
	for {
		pool, err := pgxpool.Connect(context.Background(), connString)
		if err != nil {
			log.Printf("[E] Could not connect to Postgres at <%s>: %s", connString, err)
			time.Sleep(5 * time.Second)
			continue
		}
		return pool, nil
	}
}
