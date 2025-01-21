package service

import (
	"context"

	"github.com/serj213/music-playlist/internal/app/domain"
)


type PlaylistRepository interface {
	AddSong(ctx context.Context, song domain.Song) (int64, error)
	GetSongById(ctx context.Context, id int) (domain.Song, error)
	GetSongOnPosition(ctx context.Context, position int) (domain.Song, error)
	GetPlaylist(ctx context.Context) ([]domain.Song, error)
}