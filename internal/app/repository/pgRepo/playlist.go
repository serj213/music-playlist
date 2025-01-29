package pgrepo

import (
	"context"

	"github.com/serj213/music-playlist/internal/app/domain"
)

type PlaylistRepo  struct {
	db *PgDb
}


func NewPlaylistRepo(db *PgDb) *PlaylistRepo {
	return &PlaylistRepo{
		db: db,
	}
}

func (r PlaylistRepo) CreatePlaylist(ctx context.Context, playlist domain.Playlist) (int, error) {

	// playlistModel

}