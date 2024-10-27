package dtos

type SectionEvent struct {
	// Id       int    `json:"id"`
	Name     string `json:"name" form:"section"`
	Price    int    `json:"price" form:"price"`
	Quantity int    `json:"quantity" form:"quantity"`
	EventId  int    `json:"eventId"`
}
