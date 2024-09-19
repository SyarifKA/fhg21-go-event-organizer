package models

type Categories struct {
	Id    int    `json:"id"`
	Image string `json:"image" db:"image"`
	Title string `json:"title" db:"title"`
	Date  string `json:"date" db:"date"`
}
