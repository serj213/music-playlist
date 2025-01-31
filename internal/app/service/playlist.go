package service

import (
	"context"

	"github.com/serj213/music-playlist/internal/app/domain"
)


type PlaylistService struct {
	repo playlistRepository
}

func NewPlaylistService(repo playlistRepository) PlaylistService {
	return PlaylistService{
		repo: repo,
	}
}

func (p PlaylistService) CreatePlaylist(ctx context.Context, playlist domain.Playlist) (domain.Playlist, error) {
	return p.repo.CreatePlaylist(ctx, playlist)
}