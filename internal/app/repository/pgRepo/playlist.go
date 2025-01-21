package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/serj213/music-playlist/internal/app/domain"
	"github.com/serj213/music-playlist/internal/app/repository/models"
	"github.com/serj213/music-playlist/internal/pkg/pg"
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
	songModel := domainToSong(song)

	res, err := r.db.NewInsert().Model(&songModel).Exec(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed insert song: %v", err)
	}

	songId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed get id insert song: %v", err)
	}
	return songId, nil
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