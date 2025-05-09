package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func CreateProfile(data models.Profile) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		INSERT INTO profile (full_name, user_id)
		VALUES ($1, $2) RETURNING *
		`
	row, err := db.Query(context.Background(), sql, data.FullName, data.UserId)

	if err != nil {
		return models.Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Profile])

	if err != nil {
		return models.Profile{}, nil
	}

	return profile, err
}

func ListAllProfile() []dtos.JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "u"."email", "p"."full_name", "u"."username", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id"`

	rows, _ := db.Query(
		context.Background(),
		joinSql,
	)

	events, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.JoinProfile])
	return events
}

func FindProfileByUserId(id int) models.FindProfileById {
	db := lib.DB()
	defer db.Close(context.Background())

	// joinSql := `select "u"."id", "u"."email", "u"."username", "p"."picture", "p"."full_name", "p"."birth_date", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."user_id"
	// from "users" "u"
	// left join "profile" "p"
	// on "p"."id" = "p"."user_id"
	// where "u"."id" = $1`

		joinSql := `select "u"."id",
		"u"."email",
		"u"."username",
		"p"."picture",
		"p"."full_name",
		TO_CHAR("p"."birth_date", 'YYYY-MM-DD') AS "birth_date",
		"p"."gender",
		"p"."phone_number",
		"p"."profession",
		"p"."nationality_id",
		"p"."user_id"
		FROM "users" "u"
		LEFT JOIN "profile" "p" ON "u"."id" = "p"."user_id"
		WHERE "u"."id" = $1`

	rows, _ := db.Query(
		context.Background(),
		joinSql, id,
	)
	
	profile, _ := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.FindProfileById])
	// fmt.Println(profile)
	// fmt.Println(profile)
	return profile
}

func EditProfile(data models.UpdateProfile) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	fmt.Println(data.PhoneNumber)
	// convertDate := data.BirthDate.Local().String()
	// fmt.Println(convertDate)

	sqlProfile := `update "profile" set (full_name, birth_date, phone_number, gender, profession, nationality_id) = ($1, $2, $3, $4, $5, $6) where id = $7 returning *`
	row, err := db.Query(context.Background(), sqlProfile, data.FullName, data.BirthDate, data.PhoneNumber, data.Gender, data.Profession, data.NationalityId, data.UserId)

	if err != nil {
		fmt.Println(err)
	}

	result, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Profile])

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	return result, err
}

func UpdateProfileImage(data models.Profile, id int) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET "picture" = $1 WHERE user_id=$2 returning *`

	row, err := db.Query(context.Background(), sql, data.Picture, id)
	if err != nil {
		return models.Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Profile])
	if err != nil {
		return models.Profile{}, nil
	}

	return profile, nil
}
