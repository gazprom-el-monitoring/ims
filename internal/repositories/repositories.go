package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{}
}
