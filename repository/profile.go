package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

// func CreateProfile(data Profile) Profile {
// 	db := lib.DB()
// 	defer db.Close(context.Background())
// 	row := db.QueryRow(
// 		context.Background(),
// 		`insert into "profile" (picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id) values ($1, $2, $3, $4, $5, $6, $7, $8) returning "id", "picture", "fullName", "birthDate", "gender", "phoneNumber", "profession", "nationalityId", "userId"`,
// 		data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UserId,
// 	)
// 	var results Profile
// 	row.Scan(
// 		&results.Picture,
// 		&results.FullName,
// 		&results.BirthDate,
// 		&results.Gender,
// 		&results.PhoneNumber,
// 		&results.Profession,
// 		&results.NationalityId,
// 		&results.UserId,
// 	)
// 	return results
// }

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

func FindProfileByUserId(id int) dtos.Profiles {
	db := lib.DB()
	defer db.Close(context.Background())

	// joinSql := `select * from "profile" where "user_id" = $1;`
	// joinSql := `select *
	// from "users" "u"
	// join "profile" "p"
	// where "u"."id" = $1;`

	joinSql := `select "u"."id", "u"."email", "u"."username", "p"."picture", "p"."full_name", "p"."birth_date", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."user_id"
	from "users" "u"
	join "profile" "p"
	on "u"."id" = "p"."user_id"
	where "u"."id" = $1;`

	rows, _ := db.Query(
		context.Background(),
		joinSql, id,
	)

	profile, _ := pgx.CollectOneRow(rows, pgx.RowToStructByPos[dtos.Profiles])
	// fmt.Println(profile)

	return profile
}

func EditProfile(data dtos.Profiles, id int) dtos.Profiles {
	db := lib.DB()
	defer db.Close(context.Background())

	// data.Password = lib.Encrypt(data.Password)

	// sqlUser := `update "users" set (username) values = ($1) where "id" = $2 returning "username"`
	// sqlUser := `update "users" set (full_name, username, email, phone_number) values = ($1, $2, $3, $4) where "id" = $5 returning "username"`
	// var userId int
	// user := db.QueryRow(context.Background(), sqlUser, data.Username, id)

	// fmt.Println(err1)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	fmt.Println("err1")
	// }

	fmt.Println(data.PhoneNumber)

	sqlProfile := `update "profile" set (phone_number) values = ($1) where id = $2 returning "id", "phone_number"`
	row := db.QueryRow(context.Background(), sqlProfile, data.PhoneNumber, id)

	var result dtos.Profiles
	row.Scan(
		&result.Id,
		&result.PhoneNumber,
	)
	fmt.Println(result.PhoneNumber)
	fmt.Println(result.Id)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	return result
}
