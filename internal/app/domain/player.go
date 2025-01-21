package domain


type Song struct {
	id int
	title string
	artist string
	duration int
	file_path string
	position int
}


type NewSongData struct {
	Id int
	Title string
	Artist string
	Duration int
	File_path string
	Position int
}

func NewSong(data NewSongData) Song {
	return Song{
		id: data.Id,
		title: data.Title,
		artist: data.Artist,
		duration: data.Duration,
		file_path: data.File_path,
		position: data.Position,
	}
}

func (p *Song) ID() int {
	return p.id
}

func (p *Song) Title() string {
	return p.title
}

func (p *Song) Artist() string {
	return p.artist
}

func (p *Song) Duration() int {
	return p.duration
}

func (p *Song) FilePath() string {
	return p.file_path
}

func (p *Song) Position() int {
	return p.position
}