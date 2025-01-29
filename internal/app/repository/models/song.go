package models

import (
	"time"

	"github.com/uptrace/bun"
)


type Song struct {
	bun.BaseModel `bun:"table:song"`
	Id int `bun:",pk,autoincrement"`
	Title string `bun:",notnull"`
	Artist string `bun:",notnull"`
	Duration string `bun:",notnull"`
	Position int
	Created_at time.Time `bun:",nullzero"`
	Updated_at time.Time `bun:",nullzero"`
}