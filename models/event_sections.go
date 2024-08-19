package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type SectionEvent struct {
	Id       	int    `json:"id"`
	Name     	string `json:"name"`
	Price 		int `json:"price"`
	Quantity    int    `json:"quantity"`
	EventId 	int `json:"eventId"`
}

func FindSectionEventId(EventId int) []SectionEvent {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "event_sections" where "event_id" = $1`,
		EventId,
	)
	sectionEvent, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[SectionEvent])

	return sectionEvent
}