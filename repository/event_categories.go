package repository

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func CreateEventCategories(user models.EventCategories) models.EventCategories {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "event_categories" (event_id, category_id) values ($1, $2) returning *`

	row, err := db.Query(context.Background(), sql, user.EventId, user.CategoryId)

	if err != nil {
		return models.EventCategories{}
	}

	categories, _ := pgx.CollectOneRow(row, pgx.RowToStructByName[models.EventCategories])

	return categories

}
