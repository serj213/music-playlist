package pgrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
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

// Подумать как будет песня добавляться в плейлист

/*
	При добавлении песни в плейлист, я получаю с клиента Id плейлиста
*/

func (r PgDb) AddSong(ctx context.Context, song domain.Song, playlistId int)(int, error){

	songModel := domainToSong(song)
	var songId int

	err := pg.HandlePgTransaction(ctx, func(pgTX pgx.Tx) error {
		err := r.db.QueryRow(ctx,
			"INSERT INTO song (title, artist, duration) VALUES($1,$2,$3) RETURNING id", 
			songModel.Title, songModel.Artist, songModel.Duration).Scan(&songId)
	   if err != nil {
		   return fmt.Errorf("failed insert song: %w", err)
	   }
   
		_, err = r.db.Exec(ctx, 
		   "INSERT INTO playlist_songs(playlist_id, song_id) VALUES($1,$2)", 
		   playlistId, songId)
   
	   if err != nil {
		   return fmt.Errorf("failed instert into playlist_song: %w", err)
	   }

	   return nil
	}, r.db)

	if err != nil {
		return 0, err
	}

	return songId, nil
}

// Когда использовать транзакции