package pgrepo

import (
	"context"
)

type PlaylistRepo  struct {
	db *PgDb
}


func NewPlaylistRepo(db *PgDb) *PlaylistRepo {
	return &PlaylistRepo{
		db: db,
	}
}

func (r PlaylistRepo) CreatePlaylist(ctx context.Context) {

}