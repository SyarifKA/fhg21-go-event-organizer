package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Locations struct{
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Lat string `json:"lat"`
	Long string `json:"long"`
}

func FindAllLocation() []*Locations{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "locations"`
	rows, _ := db.Query(context.Background(), sql)

	locations,_ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Locations])
	return locations
}