package dtos

type EventCategories struct {
	// Id         int `json:"id"`
	EventId    int `json:"eventId"`
	CategoryId int `json:"categoryId" form:"categoryId"`
}
