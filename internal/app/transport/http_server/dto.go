package httpserver

import "github.com/serj213/music-playlist/internal/app/domain"

func playlistToDomain(newPlaylist PlaylistRequest) (domain.Playlist, error) {
	return domain.NewPlaylist(domain.NewPlaylistData{
		Title: newPlaylist.Title,
	})
}

func songToDomain(song AddSongRequest)(domain.Song, error) {
	return domain.NewSong(domain.NewSongData{
		Title: song.Title,
		Artist: song.Artist,
		Duration: song.Duration,
	}),nil
}