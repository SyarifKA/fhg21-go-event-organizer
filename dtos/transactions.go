package dtos

type FormTransactions struct {
	EventId         int   `json:"eventId" form:"eventId" db:"event_id"`
	PaymentMethodId int   `json:"paymentMethodId" form:"paymentMethodId" db:"payment_method_id"`
	SectionId       []int `json:"sectionId" form:"sectionId[]" db:"section_id"`
	TicketQty       []int `json:"ticketQty" form:"ticketQty[]" db:"ticket_qty"`
}
