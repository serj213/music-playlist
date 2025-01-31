package models

import (
	"time"
)


type Playlist struct {
	Id int 
	Title string
	Created_at time.Time 
	Updated_at time.Time 
}