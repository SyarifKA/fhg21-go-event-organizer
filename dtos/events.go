package dtos

type Events struct {
	Id          int    `json:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title       string `json:"title" form:"title" binding:"required" db:"title"`
	Date        string `json:"date" form:"date" db:"date"`
	Description string `json:"description" form:"description" binding:"required" db:"description"`
	LocationId  *int   `json:"locationId" db:"location_id"`
	CreatedBy   *int   `json:"createdBy" db:"created_id"`
}
