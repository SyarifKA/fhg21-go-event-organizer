package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
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

func FindAllCategories(search string, limit int, page int) ([]dtos.Categories, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	inputSQL := `select * from "categories" where "name" ilike '%' || $1 || '%' limit $2 offset $3`
	rows, _ := db.Query(context.Background(), inputSQL, search, limit, offset)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Categories])
	if err != nil {
		fmt.Println(err)
	}
	count := TotalCategories(search)
	return users, count
}

func TotalCategories(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())
	inputSQL := `select count(id) as "total" from "categories" where "name" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), inputSQL, search)
	var result int
	rows.Scan(
		&result,
	)
	return result
}

func FindOneCategoriesById(id int) dtos.Categories {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "categories" where "id" = $1`,
		id,
	)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Categories])

	if err != nil {
		fmt.Println(err)
	}

	category := dtos.Categories{}
	for _, item := range categories {
		if item.Id == id {
			category = item
		}
	}
	return category
}

func CreateCategories(user dtos.Categories) dtos.Categories {
	db := lib.DB()
	defer db.Close(context.Background())

	row := db.QueryRow(
		context.Background(),
		`insert into "categories" (name) values ($1) returning "id", "name"`,
		user.Name,
	)

	var results dtos.Categories
	row.Scan(
		&results.Id,
		&results.Name,
	)
	return results
}

func DeleteCategories(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`delete from "categories" where id=$1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func EditCategories(user dtos.Categories, id int) dtos.Categories {
	db := lib.DB()
	defer db.Close(context.Background())
	// user.Password = lib.Encrypt(user.Password)

	dataSql := `update "categories" set (name) = ($1) where "id" = $2`

	db.Exec(context.Background(), dataSql, user.Name, id)
	user.Id = id
	return user
}
