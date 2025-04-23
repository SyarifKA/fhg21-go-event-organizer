package models

type Events struct {
	Id          int    `json:"id"`
	Image       string `json:"image" db:"image"`
	Title       string `json:"title" db:"title"`
	Date        string `json:"date" db:"date"`
	Description string `json:"description" db:"description"`
	LocationId  int    `json:"locationId" db:"location_id"`
	CreatedBy   int    `json:"createdBy" db:"created_id"`
}
