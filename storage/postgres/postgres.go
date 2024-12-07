package postgres

import (
	"context"
	"fmt"
	"task/config"
	"task/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db   *pgxpool.Pool
	song storage.SongRepoI
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	),
	)
	if err != nil {
		return nil, err
	}

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pgxpool,
	}, err
}

func (s *Store) Song() storage.SongRepoI {
	if s.song == nil {
		s.song = NewSongRepo(s.db)
	}
	return s.song
}
