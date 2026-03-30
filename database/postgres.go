package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(databaseURL string) (*Postgres, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(context.Background()); err != nil {
		_ = db.Close()
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.DB.PingContext(ctx)
}

func (p *Postgres) Close() {
	_ = p.DB.Close()
}
