package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id          int    `json:"id"`
	Image string `json:"image" form:"image"`
	Title       string `json:"title" form:"title" binding:"required,title" db:"title"`
	Date        string `json:"date" form:"date" binding:"required" db:"date"`
	Description string `json:"description" form:"description" binding:"required" db:"description"`
}

// "id" serial primary key,
//     "image" varchar(255),
//     "title" varchar(50),
//     "date" timestamp,
//     "description" text,
//     "location_id" int references "locations"("id"),
//     "created_by" int references "users"("id")

func FindAllEvents() []Events {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "events" order by "id" asc`)
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])
	if err != nil {
		fmt.Println(err)
	}
	return events
}

func CreateEvent(event Events) Events {
	db := lib.DB()
	defer db.Close(context.Background())
	// fmt.Println(event)

	row := db.QueryRow(
		context.Background(),
		`insert into "events" (image, title, date, description) values ($1, $2, $3, $4) returning "id", "image", "title", "date", "description"`,
		event.Image, event.Title, event.Date, event.Description,
	)
	
	var results Events
	row.Scan(
		&results.Id,
		&results.Image,
		&results.Title,
		&results.Date,
		&results.Description,
	)
	return results
}