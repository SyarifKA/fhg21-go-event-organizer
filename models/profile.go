package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

// type Profile struct{
// 	Id int `json:"id"`
// 	Picture string `json:"picture" form:"picture"`
// 	FullName string `json:"fullName" form:"fullName" db:"full_name"`
// 	BirthDate string `json:"birthDate" form:"birthDate" db:"birth_date"`
// 	Gender int `json:"gender" form:"gender" db:"gender"`
// 	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
// 	Profession string `json:"profession" form:"profession" db:"profession"`
// 	NationalityId int `json:"nationalityId" db:"nationality_id"`
// 	UserId int `json:"userId" db:"user_id"`
// }

type Profile struct {
	Id 				int `json:"id" db:"id"`
	Picture 		*string `json:"picture" db:"picture"`
	FullName 		string `json:"fullName" form:"fullName" db:"full_name"`
	BirthDate 		*string `json:"birthDate" db:"birth_date"`
	Gender 			int `json:"gender" db:"gender"`
	PhoneNumber 	*string `json:"phoneNumber" db:"phone_number"`
	Profession		*string `json:"profession" db:"profession"`
	NationalityId 	*int `json:"nationalityId" db:"nationality_id"`
	UserId 			int `json:"userId" db:"user_id"`
}

type JoinProfile struct {
	Id 				int `json:"id"`
	// FullName 		string `json:"fullName"`
	Email 			string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Profile Profile 
	// Gender 			int `json:"gender,omitempty"`
	// PhoneNumber 	*string `json:"phoneNumber,omitempty"`
	// Profession		*string `json:"profession,omitempty"`
	// BirthDate 		*string `json:"birthDate,omitempty"`
	// Nationality		int `json:"nationality,omitempty"`
}

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

func CreateProfile(data JoinProfile) *Profile {
	db := lib.DB()
	defer db.Close(context.Background())
	
	var profile Profile

	data.Password = lib.Encrypt(data.Password)

	sqlRegist := `insert into "users" 
	("email", "password") 
	values 
	($1, $2) returning "id"`
	
	var userId int
	err1 := db.QueryRow(context.Background(), sqlRegist, data.Email, data.Password).Scan(
		&userId,
	)
	fmt.Println(err1)
	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("err1")
	}

	sqlProfile := `insert into "profile" 
	("picture","full_name","birth_date","gender","phone_number","profession", "nationality_id", "user_id") 
	values 
	($1, $2, $3, $4, $5, $6, $7, $8) returning "id", "picture", "full_name", "birth_date","gender","phone_number","profession", "nationality_id", "user_id"`

	err2 := db.QueryRow(context.Background(), sqlProfile, data.Profile.Picture, data.Profile.FullName, data.Profile.BirthDate, data.Profile.Gender, data.Profile.PhoneNumber, data.Profile.Profession, data.Profile.NationalityId, userId).Scan(
		&profile.Id,
		&profile.Picture,
		&profile.FullName,
		&profile.BirthDate,
		&profile.Gender,
		&profile.PhoneNumber,
		&profile.Profession,
		&profile.NationalityId,
		&profile.UserId,
	)
	
	if err2 != nil {
		fmt.Println(err2)
	}

	// result := JoinProfile{}
	
	// result.Id = data.UserId
	// result.Password = data.Password
	// result.FullName = data.FullName
	// fmt.Println(result)
	
	return &profile
}

func ListAllProfile()[]JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "u"."email", "p"."full_name", "u"."username", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id"`
	
	rows, _:= db.Query(
		context.Background(),
		joinSql,
	)
	
	events, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[JoinProfile])
	return events
}

func FindProfileByUserId(id int) JoinProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	var result JoinProfile
	for _, v := range ListAllProfile() {
		if v.Id == id {
			result = v
		}
	}

	return result
}