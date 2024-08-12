package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id          int    `json:"id"`
	Image string `json:"image" form:"image" db:"image"`
	Title       string `json:"title" form:"title" binding:"required" db:"title"`
	Date        string `json:"date" form:"date" db:"date"`
	Description string `json:"description" form:"description" binding:"required" db:"description"`
	LocationId *int `json:"locationId" db:"location_id"`
	CreatedBy *int `json:"createdBy" db:"created_id"`
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
	fmt.Println(event)

	row := db.QueryRow(
		context.Background(),
		`insert into "events" (image, title, date, description) values ($1, $2, $3, $4) returning "id", "image", "title", "date", "description"`,
		event.Image, event.Title, event.Date, event.Description,
	)
	
	// var results Events
	results := Events{}
	fmt.Println(event.Date)
	row.Scan(
		&results.Id,
		&results.Image,
		&results.Title,
		&results.Date,
		&results.Description,
	)
	return results
}

func DeleteEvent(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`delete from "events" where id=$1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func FindOneEventById(id int) Events {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "events" where "id" = $1`,
		id,
	)
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])

	if err != nil {
		fmt.Println(err)
	}

	event := Events{}
	// fmt.Println(event)
	// fmt.Println(event)
	for _, item := range events {
		if item.Id == id {
			event = item
		}
	}
	return event
}

func EditEvent(data Events, id int) Events {
// func EditEvent(image string, title string, date string, description string, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "events" set (image , title, date, description) = ($1, $2, $3, $4) where "id" = $5`

	db.Exec(context.Background(), dataSql, data.Image, data.Title, data.Date, data.Description, id)
	data.Id = id
	return data
}
