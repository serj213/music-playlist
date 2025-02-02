package httpserver

import (
	"context"

	"github.com/serj213/music-playlist/internal/app/domain"
)


type playlistService interface {
	CreatePlaylist(ctx context.Context, playlist domain.Playlist) (domain.Playlist, error)
}

type songService interface {
	AddSong(ctx context.Context, song domain.Song, playlistId int) (int, error)
}