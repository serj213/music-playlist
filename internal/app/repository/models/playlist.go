package models

import (
	"time"

	"github.com/uptrace/bun"
)


type Playlist struct {
	bun.BaseModel `bun:"table:playlist"`
	Id int `bun:",pk,autoincrement"`
	Title string
	Created_at time.Time `bun:",nullzero"`
	Updated_at time.Time `bun:",nullzero"`
}