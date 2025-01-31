package httpserver

import "github.com/serj213/music-playlist/internal/app/domain"

func playlistToDomain(newPlaylist PlaylistRequest) (domain.Playlist, error) {
	return domain.NewPlaylist(domain.NewPlaylistData{
		Title: newPlaylist.Title,
	})
}