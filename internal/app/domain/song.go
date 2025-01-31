package domain


type Song struct {
	id int
	title string
	artist string
	duration int
}


type NewSongData struct {
	Id int
	Title string
	Artist string
	Duration int
}

func NewSong(data NewSongData) Song {
	return Song{
		id: data.Id,
		title: data.Title,
		artist: data.Artist,
		duration: data.Duration,
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

func (p Song) Duration() int {
	return p.duration
}
