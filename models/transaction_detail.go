package models

type TransactionDetail struct {
	Id            int `json:"id"`
	TransactionId int `json:"transactionId" db:"transaction_id"`
	SectionId     int `json:"sectionId" db:"section_id"`
	TicketQty     int `json:"ticketQty" db:"ticket_qty"`
}
