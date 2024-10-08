package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllLocation() []*models.Locations {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "locations"`
	rows, _ := db.Query(context.Background(), sql)

	locations, _ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.Locations])
	return locations
}
