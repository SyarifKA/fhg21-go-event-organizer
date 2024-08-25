package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Wishlist struct {
	Id      int `json:"id"`
	UserId  int `json:"userId" db:"user_id"`
	EventId int `json:"eventId" db:"event_id"`
}

func InputWishList(userId int, event int) Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	// fmt.Println(event)
	// fmt.Println(userId)
	sql := `insert into "wishlist" (user_id, event_id) values ($1, $2) returning "id", "user_id", "event_id"`

	wishlist := db.QueryRow(context.Background(), sql, userId, event)

	var result Wishlist
	wishlist.Scan(
		&result.Id,
		&result.UserId,
		&result.EventId,
	)
	fmt.Println(result)
	return result
}

func FindAllWishlist() []Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "wishlist"`

	rows, err := db.Query(context.Background(), sql)

	wishlist, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[Wishlist])

	if err != nil {
		panic(err)
	}

	return wishlist
}
