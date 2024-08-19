package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
)

func CreateWishList() {
	db := lib.DB()
	defer db.Close(context.Background())

	// sql := `insert into "wishlist" (user_id, event_id) values $1, $2 returning`
}