package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"task/api/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type songRepo struct {
	db *pgxpool.Pool
}

func NewSongRepo(db *pgxpool.Pool) *songRepo {
	return &songRepo{
		db: db,
	}
}

func (r *songRepo) Create(ctx context.Context, req *models.CreateSong) (*models.Song, error) {

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var (
		songId = uuid.New().String()
		query  = `
			INSERT INTO "songs"(
				"id",
				"group_name",
				"song_name",
				"release_date",
				"lyrics",
				"link",
				"created_at"
			) VALUES($1, $2, $3, $4, $5, $6, NOW())
		`
	)

	_, err = r.db.Exec(ctx, query,
		songId,
		req.GroupName,
		req.SongName,
		req.ReleaseDate,
		req.Lyrics,
		req.Link,
	)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	queryGenre := `INSERT INTO "song_genres"("song_id", "genre") VALUES($1, $2)`

	for _, genre := range req.Genres {
		_, err := tx.Exec(ctx, queryGenre, songId, genre)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return r.GetById(ctx, &models.SongPrimaryKey{Id: songId})
}

func (r *songRepo) GetById(ctx context.Context, req *models.SongPrimaryKey) (*models.Song, error) {
	var (
		query = `
            SELECT 
                "id",
                "group_name",
                "song_name",
                "release_date",
                "lyrics",
                "link",
                "created_at",
                "updated_at"
            FROM "songs"
            WHERE id = $1
        `
		queryGenre = `
            SELECT 
                "genre"
            FROM "song_genres"
            WHERE "song_id" = $1
        `
	)

	var (
		id          sql.NullString
		groupName   sql.NullString
		songName    sql.NullString
		releaseDate sql.NullString
		lyrics      sql.NullString
		link        sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
		genre       sql.NullString
		genres      = []string{}
	)

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&groupName,
		&songName,
		&releaseDate,
		&lyrics,
		&link,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch song: %w", err)
	}

	rows, err := r.db.Query(ctx, queryGenre, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch genres: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&genre); err != nil {
			return nil, fmt.Errorf("failed to scan genre: %w", err)
		}
		if genre.Valid {
			genres = append(genres, genre.String)
		}
	}

	return &models.Song{
		Id:          id.String,
		GroupName:   groupName.String,
		SongName:    songName.String,
		ReleaseDate: releaseDate.String,
		Lyrics:      lyrics.String,
		Link:        link.String,
		Genres:      genres,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *songRepo) Delete(ctx context.Context, req *models.SongPrimaryKey) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, "DELETE FROM song_genres WHERE song_id = $1", req.Id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	_, err = r.db.Exec(ctx, "DELETE FROM songs WHERE id = $1", req.Id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *songRepo) GetAll(ctx context.Context, req *models.GetSongRequest) (*models.GetSongResponse, error) {
	var (
		query = `
			SELECT 
				"id",
				"group_name",
				"song_name",
				"release_date",
				"lyrics",
				"link",
				"created_at",
				"updated_at"
			FROM "songs"
			`

		queryGenre = `
			SELECT 
				"genre"
			FROM "song_genres"
			WHERE "song_id" = $1`

		resp   models.GetSongResponse
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		sort   = " ORDER BY created_at DESC"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if len(req.Search) > 0 {
		where += " AND group_name ILIKE" + " '%" + req.Search + "%'" + " OR song_name ILIKE" + " '%" + req.Search + "%'"
	}

	query += where + sort + offset + limit
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			song        models.Song
			id          sql.NullString
			groupName   sql.NullString
			songName    sql.NullString
			releaseDate sql.NullString
			lyrics      sql.NullString
			link        sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
			genre       sql.NullString
			genres      = []string{}
		)

		err = rows.Scan(
			&id,
			&groupName,
			&songName,
			&releaseDate,
			&lyrics,
			&link,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		rowsGenres, err := r.db.Query(ctx, queryGenre, id.String)
		if err != nil {
			return nil, err
		}
		for rowsGenres.Next() {
			err = rowsGenres.Scan(
				&genre,
			)
			if err != nil {
				return nil, err
			}
			genres = append(genres, genre.String)
		}

		song = models.Song{
			Id:          id.String,
			GroupName:   groupName.String,
			SongName:    songName.String,
			ReleaseDate: releaseDate.String,
			Lyrics:      lyrics.String,
			Link:        link.String,
			Genres:      genres,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		}
		resp.Songs = append(resp.Songs, &song)
	}
	return &resp, nil
}
func (r *songRepo) Update(ctx context.Context, req *models.UpdateSong) (int64, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	deleteQuery := `DELETE FROM "song_genres" WHERE song_id = $1`
	_, err = tx.Exec(ctx, deleteQuery, req.Id)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	updateQuery := `
		UPDATE "songs"
		SET
			group_name = $1,
			song_name = $2,
			release_date = $3,
			lyrics = $4,
			link = $5,
			updated_at = NOW()
		WHERE id = $6
	`
	result, err := tx.Exec(ctx, updateQuery,
		req.GroupName,
		req.SongName,
		req.ReleaseDate,
		req.Lyrics,
		req.Link,
		req.Id,
	)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	queryGenre := `INSERT INTO "song_genres"("song_id", "genre") VALUES($1, $2)`

	for _, genre := range req.Genres {
		_, err := tx.Exec(ctx, queryGenre, req.Id, genre)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	return rowsAffected, nil
}
