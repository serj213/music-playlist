package models

import "github.com/uptrace/bun"


type Playlist struct {
	bun.BaseModel `bun:"table:playlist"`
	Id int
	Title string
	Artist string
	Duration int
	File_path string
}