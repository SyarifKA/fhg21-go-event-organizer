package dtos

type Events struct {
	Id          int    `json:"id"`
	Title       string `json:"title" form:"title" binding:"required"`
	Date        string `json:"date" form:"date"`
	Description string `json:"description" form:"description" binding:"required"`
	LocationId  *int   `json:"locationId" form:"locationId" db:"location_id"`
	CreatedBy   *int   `json:"createdBy"`
	CategoryId  int    `json:"categoryId" form:"categoryId"`
}

type ImageEvent struct {
	Image string `json:"image" form:"eventImg"`
}
