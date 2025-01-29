package domain


type Song struct {
	id int
	title string
	artist string
	duration string
	position int
}


type NewSongData struct {
	Id int
	Title string
	Artist string
	Duration string
	Position int
}

func NewSong(data NewSongData) Song {
	return Song{
		id: data.Id,
		title: data.Title,
		artist: data.Artist,
		duration: data.Duration,
		position: data.Position,
	}
}

func (p Song) ID() int {
	return p.id
}

func (p Song) Title() string {
	return p.title
}

func (p Song) Artist() string {
	return p.artist
}

func (p Song) Duration() string {
	return p.duration
}

func (p Song) Position() int {
	return p.position
}