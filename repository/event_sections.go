package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindSectionEventId(EventId int) []models.SectionEvent {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "event_sections" where "event_id" = $1`,
		EventId,
	)
	sectionEvent, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.SectionEvent])

	return sectionEvent
}

func CreateEventSection(event models.SectionEvent) models.SectionEvent {
	db := lib.DB()
	defer db.Close(context.Background())
	sql := `INSERT into "event_sections" (name, price, quantity, event_id) values ($1, $2, $3, $4) returning *`
	row, err := db.Query(context.Background(), sql, event.Name, event.Price, event.Quantity, event.EventId)

	if err != nil {
		return models.SectionEvent{}
	}
	section, _ := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.SectionEvent])
	// fmt.Println(section)
	return section
}
