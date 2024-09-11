package dtos

type Events struct {
	Id          int    `json:"id"`
	Image       string `json:"image" form:"image"`
	Title       string `json:"title" form:"title" binding:"required"`
	Date        string `json:"date" form:"date"`
	Description string `json:"description" form:"description" binding:"required"`
	LocationId  *int   `json:"locationId" form:"locationId" db:"location_id"`
	CreatedBy   *int   `json:"createdBy"`
}
