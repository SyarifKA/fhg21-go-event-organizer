package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type PaymentMethod struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	// Quantity string `json:"quantity"`
	// Price    int    `json:"price"`
	// EventId int `json:"eventId"`
}

func FindAllPaymentMethod(method PaymentMethod) []PaymentMethod {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "payment_methods"`,
	)
	sectionEvent, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[PaymentMethod])

	return sectionEvent
}