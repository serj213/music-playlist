package models


type Song struct {
	Id int `db:"id"`
	Title string `db:"title"`
	Artist string `db:"artist"`
	Duration int `db:"duration"`
}