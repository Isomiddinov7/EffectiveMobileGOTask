package storage

import (
	"context"
	"task/api/models"
)

type StorageI interface {
	Song() SongRepoI
}

type SongRepoI interface {
	Create(ctx context.Context, req *models.CreateSong) (*models.Song, error)
	GetById(ctx context.Context, req *models.SongPrimaryKey) (*models.Song, error)
	GetAll(ctx context.Context, req *models.GetSongRequest) (*models.GetSongResponse, error)
	Update(ctx context.Context, req *models.UpdateSong) (int64, error)
	Delete(ctx context.Context, req *models.SongPrimaryKey) error
}
