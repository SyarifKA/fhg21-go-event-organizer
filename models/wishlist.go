package models

type Wishlist struct {
	Id      int `json:"id"`
	UserId  int `json:"userId" db:"user_id"`
	EventId int `json:"eventId" db:"event_id"`
}
