package pgrepo

import (
	"github.com/serj213/music-playlist/internal/app/domain"
	"github.com/serj213/music-playlist/internal/app/repository/models"
)



func domainToSong(song domain.Song) *models.Playlist {
	return &models.Playlist{
		Id: song.ID(),
		Title: song.Title(),
		Artist: song.Artist(),
		Duration: song.Duration(),
	}
}

func songToDomain(data models.Playlist) (domain.Song, error) {
	return domain.NewSong(domain.NewSongData{
		Id: data.Id,
		Title: data.Title,
		Artist: data.Artist,
		Duration: data.Duration,
	}),nil
}