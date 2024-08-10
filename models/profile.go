package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Profile struct{
	Id int `json:"id"`
	Picture string `json:"picture" form:"picture"`
	FullName string `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate string `json:"birthDate" form:"birthDate" db:"birth_date"`
	Gender int `json:"gender" form:"gender" db:"gender"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession string `json:"profession" form:"profession" db:"profession"`
	NationalityId int `json:"nationalityId" db:"nationality_id"`
	UserId int `json:"userId" db:"user_id"`
}

func CreateProfile(data Profile) Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	row := db.QueryRow(
		context.Background(),
		`insert into "profile" (picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id) values ($1, $2, $3, $4, $5, $6, $7, $8) returning "id", "picture", "fullName", "birthDate", "gender", "phoneNumber", "profession", "nationalityId", "userId"`,
		data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UserId,
	)
	
	var results Profile
	row.Scan(
		&results.Picture,
		&results.FullName,
		&results.BirthDate,
		&results.Gender,
		&results.PhoneNumber,
		&results.Profession,
		&results.NationalityId,
		&results.UserId,
	)
	return results
}

func FindProfileByUserId(id int) Profile{
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select "p"."id", "p"."full_name", "p"."email" from "profile" "p" left join "users" "u" on "p"."user_id" = "u"."id" where "id"=$1`,
		id,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])

	if err != nil {
		fmt.Println(err)
	}

	user := Profile{}
	for _, item := range users {
		if item.Id == id {
			user = item
		}
	}
	return user
}