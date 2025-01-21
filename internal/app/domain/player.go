package domain


type Player struct {
	id int
	title string
	artist string
	duration int
	file_path string
}


type NewPlayerData struct {
	Id int
	Title string
	Artist string
	Duration int
	File_path string
}

func NewPlayer(data NewPlayerData) *Player {
	return &Player{
		id: data.Id,
		title: data.Title,
		artist: data.Artist,
		duration: data.Duration,
		file_path: data.File_path,
	}
}

func (p *Player) ID() int {
	return p.id
}

func (p *Player) Title() string {
	return p.title
}

func (p *Player) Artist() string {
	return p.artist
}

func (p *Player) Duration() int {
	return p.duration
}

func (p *Player) FilePath() string {
	return p.file_path
}