package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/serj213/music-playlist/internal/app/domain"
	"github.com/serj213/music-playlist/internal/app/repository/models"
	"github.com/serj213/music-playlist/internal/pkg/pg"
	"github.com/uptrace/bun"
)



type PgDb struct {
	db *pg.DB
}

func NewPgRepo(db *pg.DB) *PgDb {
	return &PgDb{
		db: db,
	}
}


func (r PgDb) AddSong(ctx context.Context, song domain.Song) (int64, error) {
	var result int64

	err := pg.HandleBunTransaction(ctx, func(tx bun.Tx) error {

		position, err := r.GetLastPostion(ctx)

		if err != nil {
			return err
		}

		songModel := domainToSong(song)

		res, err := r.db.NewInsert().Model(&songModel).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed insert song: %w", err)
		}

		songId, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed get insert id: %w", err)
		}

		if err := r.db.NewUpdate().Model((*models.Playlist)(nil)).Set("position = ?", position).Where("title = ?", song.Title()).Scan(ctx); err != nil {
			return fmt.Errorf("failed update position: %w", err)
		}

		result = songId

		return nil
		}, r.db) 

	if err != nil {
		return int64(0), err
	}

	return result, nil
}

func (r PgDb) GetSongById(ctx context.Context, id int) (domain.Song, error) {
	var songModel models.Playlist

	if err := r.db.NewSelect().Model(&songModel).Where("id = ?", id).Scan(ctx); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return domain.Song{}, domain.ErrNotFound
		}

		return domain.Song{}, fmt.Errorf("failed select song by id: %v", err)
	}

	song, err := songToDomain(songModel)

	if err != nil {
		return domain.Song{}, fmt.Errorf("failed convert song to domain: %v", err)
	}

	return song, nil
}

func (r PgDb) GetSongOnPosition(ctx context.Context, position int) (domain.Song, error) {

	var songModel models.Playlist

	if err := r.db.NewSelect().Model(&songModel).Where("position = ?", position).Scan(ctx); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return domain.Song{}, domain.ErrNotFound
		}

		return domain.Song{}, fmt.Errorf("failed select next song: %v", err)
	}

	song, err := songToDomain(songModel)
	
	if err != nil {
		return domain.Song{}, fmt.Errorf("failed convert song to domain: %v", err)
	}

	return song, nil
}


func (r PgDb) GetPlaylist(ctx context.Context) ([]domain.Song, error){

	var playlistModel []models.Playlist

	if  err := r.db.NewSelect().Model(&playlistModel).Scan(ctx); err != nil {
		return []domain.Song{}, fmt.Errorf("failed select playlist: %v", err)
	}

	res := make([]domain.Song, len(playlistModel))

	for idx, value := range playlistModel {

		song, err := songToDomain(value)
		if err != nil {
			return []domain.Song{}, fmt.Errorf("failed convert model to playlist: %v", err)
		}

		res[idx] = song
	}


	return res, nil
}

func (r PgDb) GetLastPostion(ctx context.Context) (int, error) {

	var position int
	
	if err := r.db.NewSelect().Model(&position).Column("position").Order("position DESC").Scan(ctx); err != nil {
		return 0, fmt.Errorf("failed select last position: %w", err)
	}

	return position, nil
}