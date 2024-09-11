package dtos

type FormTransactions struct {
	EventId         int   `json:"eventId" form:"eventId"`
	PaymentMethodId int   `json:"paymentMethodId" form:"paymentMethodId"`
	SectionId       []int `json:"sectionId" form:"sectionId[]"`
	TicketQty       []int `json:"ticketQty" form:"ticketQty[]"`
}
