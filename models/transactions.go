package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Transactions struct{
	Id 					int `json:"id"`
	EventId         	int `json:"eventId" db:"event_id"`
	PaymentMethodId 	int `json:"paymentMethodId" db:"payment_method_id"`
	// SectionId       	[]int `json:"sectionId" form:"sectionId" db:"section_id"`
	// TicketQty       	[]int `json:"ticketQty" form:"ticketQty" db:"ticket_qty"`
	UserId 				int `json:"userId" db:"user_id"`
}

// type FormTransactions struct {
// 	Id              int   `json:"id"`
// 	// EventId         int   `json:"eventId" form:"eventId" db:"event_id"`
// 	// PaymentMethodId int   `json:"paymentMethodId" form:"paymentMethodId" db:"payment_method_id"`
// }

type DetailTransaction struct{
	TransactionId 		int `json:"transactionId" db:"transaction_id"`
	FullName 			string `json:"fullName" db:"full_name"`
	Title 				string `json:"title" db:"event_title"`
	LocationId 			*int `json:"locationId" db:"location_id"`
	Date 				string `json:"date"`
	PaymentMethodId 	string `json:"paymentMethodId" db:"payment_method"`
	TicketSection 		[]string `json:"ticketSection" db:"ticket_section"`
	TicketQty 			[]int `json:"ticketQty" db:"quantity"`
}

func CreateTransaction(transaction Transactions) Transactions {
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

	results := Transactions{}
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


func DetailTransactions(id int)DetailTransaction{
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

	details := DetailTransaction{}
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

func FindOneTransactionById(id int) Transactions {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "transactions" where "id" = $1`,
		id,
	)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Transactions])

	if err != nil {
		fmt.Println(err)
	}

	category := Transactions{}
	for _, item := range categories {
		if item.Id == id {
			category = item
		}
	}
	return category
}

func FindTransactionByUserId(id int)[]DetailTransaction{
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
    where "u"."id" = $1
    group by  "t"."id", "p"."full_name", "e"."title", "e"."location_id", "e"."date", "pm"."name";`

	rows,_ := db.Query(
		context.Background(),
		inputSQL, id,
	)

	sectionEvent, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[DetailTransaction])
	// details := []DetailTransaction{}
	// rows.Scan(
	// 	// &details.Id,
	// 	&details.TransactionId,
	// 	&details.FullName,
	// 	&details.Title,
	// 	&details.LocationId,
	// 	&details.Date,
	// 	&details.PaymentMethodId,
	// 	&details.TicketSection,
	// 	&details.TicketQty,
	// )
	fmt.Println(sectionEvent)
	// transaction, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[DetailTransaction])
	return sectionEvent
}