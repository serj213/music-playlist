package pgrepo

import (
	"github.com/serj213/music-playlist/internal/app/domain"
	"github.com/serj213/music-playlist/internal/app/repository/models"
)



func domainToPlaylist(playlist domain.Playlist) *models.Playlist {
	return &models.Playlist{
		Title: playlist.Title(),
	}
}

func domainToSong(song domain.Song) *models.Song {
	return &models.Song{
		Id: song.ID(),
		Title: song.Title(),
		Artist: song.Artist(),
		Duration: song.Duration(),

	}
}

func songToDomain(data models.Song) (domain.Song, error) {
	return domain.NewSong(domain.NewSongData{
		Id: data.Id,
		Title: data.Title,
		Artist: data.Artist,
		Duration: data.Duration,
	}),nil
}


func playlistToDomain(data models.Playlist) (domain.Playlist, error) {
	return domain.NewPlaylist(domain.NewPlaylistData{
		Title: data.Title,
	})
}