package models

type SectionEvent struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	Quantity int    `json:"quantity" form:"quantity"`
	EventId  int    `json:"eventId" db:"event_id"`
}
