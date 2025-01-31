package httpserver

import "fmt"

type SongRequest struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
	Duration int `json:"duration"`
}

type PlaylistRequest struct {
	Title string `json:"title"`
}

type playlistRes struct {
	Id int `json:"id"`
	Title string `json:"title"`
}


type PlaylistResponse struct {
	Status string `json:"status"`
	Playlist playlistRes `json:"playlist"`
}

func (s *SongRequest) Validate() error {
	if s.Title == ""{
		return fmt.Errorf("title is empty")
	}
	if s.Artist == "" {
		return fmt.Errorf("artist is empty")
	}
	if s.Duration == 0 {
		return fmt.Errorf("duration is empty")
	}
	return nil
}