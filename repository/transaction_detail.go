package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
)

func CreateTransactionDetail(data models.TransactionDetail) models.TransactionDetail {
	db := lib.DB()
	defer db.Close(context.Background())

	inputSQL := `insert into "transaction_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
	row := db.QueryRow(context.Background(), inputSQL, data.TransactionId, data.SectionId, data.TicketQty)

	var detail models.TransactionDetail

	row.Scan(
		&detail.Id,
		&detail.TransactionId,
		&detail.SectionId,
		&detail.TicketQty,
	)
	return detail
}
