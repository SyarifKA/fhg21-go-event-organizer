package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllParners() []*models.Partners {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "partners"`
	rows, _ := db.Query(context.Background(), sql)

	partners, _ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.Partners])
	return partners
}
