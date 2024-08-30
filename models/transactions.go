package models

type Transactions struct {
	Id              int `json:"id"`
	EventId         int `json:"eventId" db:"event_id"`
	PaymentMethodId int `json:"paymentMethodId" db:"payment_method_id"`
	// SectionId       	[]int `json:"sectionId" form:"sectionId" db:"section_id"`
	// TicketQty       	[]int `json:"ticketQty" form:"ticketQty" db:"ticket_qty"`
	UserId int `json:"userId" db:"user_id"`
}

// type FormTransactions struct {
// 	Id              int   `json:"id"`
// 	// EventId         int   `json:"eventId" form:"eventId" db:"event_id"`
// 	// PaymentMethodId int   `json:"paymentMethodId" form:"paymentMethodId" db:"payment_method_id"`
// }

type DetailTransaction struct {
	TransactionId   int      `json:"transactionId" db:"transaction_id"`
	FullName        string   `json:"fullName" db:"full_name"`
	Title           string   `json:"title" db:"event_title"`
	LocationId      *int     `json:"locationId" db:"location_id"`
	Date            string   `json:"date"`
	PaymentMethodId string   `json:"paymentMethodId" db:"payment_method"`
	TicketSection   []string `json:"ticketSection" db:"ticket_section"`
	TicketQty       []int    `json:"ticketQty" db:"quantity"`
}
