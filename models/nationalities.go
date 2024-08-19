package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Nationalities struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func FindAllNationality() []*Nationalities{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "nationalities"`
	rows, _ := db.Query(context.Background(), sql)

	nationalities,_ := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Nationalities])
	return nationalities
}