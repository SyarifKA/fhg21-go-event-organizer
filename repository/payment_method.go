package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllPaymentMethod(method models.PaymentMethod) []models.PaymentMethod {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "payment_methods"`)
	sectionEvent, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.PaymentMethod])

	return sectionEvent
}
