package domain




type Playlist struct {
	id int
	title string
}

type NewPlaylistData struct {
	Id int
	Title string
}

func NewPlaylist(data NewPlaylistData) (Playlist, error) {
	return Playlist{
		id: data.Id,
		title: data.Title,
	}, nil
}

func (p Playlist) ID() int {
	return p.id
}

func (p Playlist) Title() string {
	return p.title
}
