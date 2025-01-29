package domain

import "time"




type Playlist struct {
	id int
	title string
	createdAt time.Time
	updatedAt time.Time
}

type NewPlaylistData struct {
	ID int
	Title string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPlaylist(data NewPlaylistData) (Playlist, error) {
	return Playlist{
		id: data.ID,
		title: data.Title,
		updatedAt: data.UpdatedAt,
		createdAt: data.CreatedAt,
	}, nil
}

func (p Playlist) ID() int {
	return p.id
}

func (p Playlist) Title() string {
	return p.title
}

func (p Playlist) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p Playlist) CreatedAt() time.Time {
	return p.createdAt
}