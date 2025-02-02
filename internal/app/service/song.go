package service

import (
	"context"
	"log/slog"

	"github.com/serj213/music-playlist/internal/app/domain"
)

type SongService struct {
	repo songRepository
	logger *slog.Logger
}


func NewSongService(log *slog.Logger, repo songRepository) SongService {

	log = log.With(slog.String("service", "song"))

	return SongService{
		repo: repo,
		logger: log,
	}
}

func (s SongService) AddSong(ctx context.Context, song domain.Song, playlistId int)(int, error) {
	return s.repo.AddSong(ctx, song, playlistId)
}