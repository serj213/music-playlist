package pgrepo

import (
	"context"
	"fmt"

	"github.com/serj213/music-playlist/internal/app/domain"
	"github.com/serj213/music-playlist/internal/app/repository/models"
	"github.com/serj213/music-playlist/internal/pkg/pg"
)

type PlaylistRepo  struct {
	db *pg.DB
}


func NewPlaylistRepo(db *pg.DB) *PlaylistRepo {
	return &PlaylistRepo{
		db: db,
	}
}

func (r PlaylistRepo) CreatePlaylist(ctx context.Context, playlist domain.Playlist) (domain.Playlist, error) {
	playlistModel := domainToPlaylist(playlist)

	rows, err := r.db.Query(ctx, 
		"INSERT INTO playlist(title) VALUES ($1) RETURNING id, title", 
		playlistModel.Title)

	if err != nil {
		return domain.Playlist{}, fmt.Errorf("failed insert playlist: %w", err)
	}

	var resModelPlaylist models.Playlist

	for rows.Next() {
		if err := rows.Scan(&resModelPlaylist.Id, &resModelPlaylist.Title); err != nil {
			return domain.Playlist{}, fmt.Errorf("failed scan rows returning: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return domain.Playlist{}, fmt.Errorf("rows err: %w", err)
	}

	resDomainPlaylist, err := playlistToDomain(resModelPlaylist)

	if err != nil {
		return domain.Playlist{}, fmt.Errorf("failed convert playlist domain: %w", err)
	}

	return resDomainPlaylist, nil
}