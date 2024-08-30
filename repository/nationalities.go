package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllNationality() []*models.Nationalities {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "nationalities"`
	rows, _ := db.Query(context.Background(), sql)

	nationalities, _ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.Nationalities])
	return nationalities
}
