package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Partners struct{
	Id int `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

func FindAllParners() []*Partners{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "partners"`
	rows, _ := db.Query(context.Background(), sql)

	partners,_ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Partners])
	return partners
}