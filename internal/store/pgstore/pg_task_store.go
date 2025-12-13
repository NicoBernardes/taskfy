package pgstore

import "github.com/jackc/pgx/v5/pgxpool"

type PgTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPgTaskStore(pool *pgxpool.Pool) PgTaskStore {
	return PgTaskStore{Queries: New(pool), Pool: pool}
}
