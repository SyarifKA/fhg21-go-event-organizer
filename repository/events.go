package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllEvents(search string, limit int, page int) ([]models.Events, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	inputSQL := `select * from "events" where "title" ilike '%' || $1 || '%' limit $2 offset $3`
	rows, _ := db.Query(context.Background(), inputSQL, search, limit, offset)
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Events])
	if err != nil {
		fmt.Println(err)
	}
	count := TotalEvents(search)
	return events, count
}

func TotalEvents(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())
	inputSQL := `select count(id) as "total" from "events" where "title" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), inputSQL, search)
	var result int
	rows.Scan(
		&result,
	)
	return result
}

func CreateEvent(event models.Events) (models.Events, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(event)

	row, _ := db.Query(
		context.Background(),
		`insert into "events" (image, title, date, description, location_id, created_by) values ($1, $2, $3, $4, $5, $6) returning *`,
		event.Image, event.Title, event.Date, event.Description, event.LocationId, event.CreatedBy,
	)

	result, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Events])

	if err != nil {
		fmt.Println(err)
	}

	return result, err
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

func FindOneEventById(id int) models.Events {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "events" where "id" = $1`,
		id,
	)
	events, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Events])

	if err != nil {
		fmt.Println(err)
	}

	return events
}

func EditEvent(data dtos.Events, id int) dtos.Events {
	// func EditEvent(image string, title string, date string, description string, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "events" set (image , title, date, description) = ($1, $2, $3, $4) where "id" = $5`

	db.Exec(context.Background(), dataSql, data.Title, data.Date, data.Description, id)
	data.Id = id
	return data
}
