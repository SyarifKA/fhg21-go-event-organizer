package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindEventByCategoryId(categories int) ([]models.Categories, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select "ec"."id", "e"."image", "e"."title", "e"."date"
	from "event_categories" "ec"
	join "events" "e" on "e"."id" = "ec"."event_id"
	where "ec"."category_id" = $1`

	rows, err := db.Query(context.Background(), sql, categories)

	if err != nil {
		return []models.Categories{}, err
	}

	result, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Categories])

	if err != nil {
		return []models.Categories{}, err
	}
	fmt.Println(result)
	return result, err
}
