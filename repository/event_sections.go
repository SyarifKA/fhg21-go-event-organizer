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
