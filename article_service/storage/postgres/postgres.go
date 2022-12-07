package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Postgres ...
type Postgres struct {
	db *sqlx.DB
}

// InitDB ...
func InitDB(psqlConfig string) (*Postgres, error) {
	var err error

	tempDB, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: tempDB,
	}, nil
}
