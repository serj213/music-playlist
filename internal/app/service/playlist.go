package service

import (
	"context"

	"github.com/serj213/music-playlist/internal/app/domain"
)


type PlaylistService struct {
	repo PlaylistRepository
}

func NewPlaylistService(repo PlaylistRepository) PlaylistService {
	return PlaylistService{
		repo: repo,
	}
}

func (p PlaylistService) AddSong(ctx context.Context, song domain.Song) (int64, error) {
	return p.repo.AddSong(ctx, song)
}


func (p PlaylistService) GetSongById(ctx context.Context, id int) (domain.Song, error) {
	return p.repo.GetSongById(ctx, id)
}

func (p PlaylistService) GetSongOnPosition(ctx context.Context, position int) (domain.Song, error) {
	return p.repo.GetSongOnPosition(ctx, position)
}

func (p PlaylistService) GetPlaylist(ctx context.Context) ([]domain.Song, error) {
	return p.repo.GetPlaylist(ctx)
}