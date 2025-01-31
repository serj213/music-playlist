package pgrepo

import (
	"context"
	"fmt"

	"github.com/serj213/music-playlist/internal/app/domain"
	"github.com/serj213/music-playlist/internal/pkg/pg"
)



type PgDb struct {
	db *pg.DB
}

func NewSongRepo(db *pg.DB) *PgDb {
	return &PgDb{
		db: db,
	}
}

func (r PgDb) AddSong(ctx context.Context, song domain.Song)(int, error){

	songModel := domainToSong(song)
	var song_id int

	err := r.db.QueryRow(ctx,
		 "INSERT INTO song (title, artist, duration) VALUES($1,$2,$3) RETURNING id", 
		 songModel.Title, songModel.Artist, songModel.Duration).Scan(&song_id)
	if err != nil {
		return 0, fmt.Errorf("failed insert song: %w", err)
	}




	return song_id, nil
}




// Когда использовать транзакции