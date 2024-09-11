package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func CreateTransaction(transaction models.Transactions) models.Transactions {
	db := lib.DB()
	defer db.Close(context.Background())
	// fmt.Println(event)
	//  fmt.Println(transaction)
	inputSQL := `insert into "transactions" (event_id, payment_method_id, user_id) values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	row := db.QueryRow(
		context.Background(),
		inputSQL,
		transaction.EventId, transaction.PaymentMethodId, transaction.UserId,
	)

	results := models.Transactions{}
	row.Scan(
		&results.Id,
		&results.EventId,
		&results.PaymentMethodId,
		&results.UserId,
	)
	fmt.Println(results)
	return results
}

// func CreateTransactionQty(transaction Transactions) Transactions {
// 	db := lib.DB()
// 	defer db.Close(context.Background())
// 	// fmt.Println(event)

// 	inputSQL := `insert into "transactions" (section_id, ticket_qty) values ($1, $2) returning "id", "section_id", "ticket_qty"`
// 	row := db.QueryRow(
// 		context.Background(),
// 		inputSQL,
// 		transaction.SectionId, transaction.TicketQty,
// 	)

// 	results := Transactions{}
// 	row.Scan(
// 		&results.Id,
// 		&results.SectionId,
// 		&results.TicketQty,
// 	)
// 	return results
// }

func DetailTransactions(id int) models.DetailTransaction {
	db := lib.DB()
	defer db.Close(context.Background())

	inputSQL := `select "t"."id", "p"."full_name", "e"."title" as "event_title", "e"."location_id", "e"."date", "pm"."name" as "payment_method", array_agg("es"."name") as "ticket_section", array_agg("td"."ticket_qty") as "quantity"
    from "transactions" "t" 
    join "users" "u" on "u"."id" = "t"."user_id"
    join "profile" "p" on "p"."user_id" = "u"."id"
    join "events" "e" on "e"."id" = "t"."event_id"
    join "payment_methods" "pm" on "pm"."id" = "t"."payment_method_id"
    join "transaction_details" "td" on "td"."transaction_id" = "t"."id"
    join "event_sections" "es" on "es"."id" = "td"."section_id"
    where "t"."id" = $1
    group by  "t"."id", "p"."full_name", "e"."title", "e"."location_id", "e"."date", "pm"."name";`

	rows := db.QueryRow(
		context.Background(),
		inputSQL, id,
	)

	details := models.DetailTransaction{}
	rows.Scan(
		// &details.Id,
		&details.TransactionId,
		&details.FullName,
		&details.Title,
		&details.LocationId,
		&details.Date,
		&details.PaymentMethodId,
		&details.TicketSection,
		&details.TicketQty,
	)
	// fmt.Println(details)
	// transaction, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[DetailTransaction])
	return details
}

func FindOneTransactionById(id int) models.Transactions {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "transactions" where "id" = $1`,
		id,
	)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Transactions])

	if err != nil {
		fmt.Println(err)
	}

	category := models.Transactions{}
	for _, item := range categories {
		if item.Id == id {
			category = item
		}
	}
	return category
}

func FindTransactionByUserId(id int) []models.DetailTransaction {
	db := lib.DB()
	defer db.Close(context.Background())

	inputSQL := `select "t"."id", "p"."full_name", "e"."title" as "event_title", "e"."location_id", "e"."date", "pm"."name" as "payment_method", array_agg("es"."name") as "ticket_section", array_agg("td"."ticket_qty") as "quantity"
    from "transactions" "t" 
    join "users" "u" on "u"."id" = "t"."user_id"
    join "profile" "p" on "p"."user_id" = "u"."id"
    join "events" "e" on "e"."id" = "t"."event_id"
    join "payment_methods" "pm" on "pm"."id" = "t"."payment_method_id"
    join "transaction_details" "td" on "td"."transaction_id" = "t"."id"
    join "event_sections" "es" on "es"."id" = "td"."section_id"
    where "t"."user_id" = $1
    group by  "t"."id", "p"."full_name", "e"."title", "e"."location_id", "e"."date", "pm"."name";`

	rows, err := db.Query(
		context.Background(),
		inputSQL, id,
	)

	if err != nil {
		fmt.Errorf("Failed to get data")
	}

	sectionEvent, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.DetailTransaction])
	if err != nil {
		fmt.Errorf("Failed to get data")
	}
	fmt.Println(sectionEvent)
	return sectionEvent
}
