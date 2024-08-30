package models

type SectionEvent struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	EventId  int    `json:"eventId"`
}
