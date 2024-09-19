package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func InputWishList(userId int, event int) models.Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "wishlist" (user_id, event_id) values ($1, $2) returning "id", "user_id", "event_id"`

	wishlist := db.QueryRow(context.Background(), sql, userId, event)

	var result models.Wishlist
	wishlist.Scan(
		&result.Id,
		&result.UserId,
		&result.EventId,
	)
	fmt.Println(result)
	return result
}

func FindAllWishlist() []models.Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select * from "wishlist"`

	rows, err := db.Query(context.Background(), sql)

	wishlist, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Wishlist])

	if err != nil {
		panic(err)
	}

	return wishlist
}
