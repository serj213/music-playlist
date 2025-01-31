package domain




type Playlist struct {
	id int
	title string
}

type NewPlaylistData struct {
	Title string
}

func NewPlaylist(data NewPlaylistData) (Playlist, error) {
	return Playlist{
		title: data.Title,
	}, nil
}

func (p Playlist) ID() int {
	return p.id
}

func (p Playlist) Title() string {
	return p.title
}
