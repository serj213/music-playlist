package httpserver

import "fmt"

type SongRequest struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
	Duration string `json:"duration"`
}

func (s *SongRequest) Validate() error {
	if s.Title == ""{
		return fmt.Errorf("title is empty")
	}
	if s.Artist == "" {
		return fmt.Errorf("artist is empty")
	}
	if s.Duration == "" {
		return fmt.Errorf("duration is empty")
	}
	return nil
}