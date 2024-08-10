package models

import (
	"context"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
)

type FormRegister struct {
	Id int `json:"id"`
	FullName        string `json:"fullName" form:"fullName" db:"full_name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"-" form:"password"`
	ConfirmPassword string `json:"-" form:"confirmPassword" binding:"eqfield=password"`
}

func RegisterUser(user FormRegister) FormRegister {
	db := lib.DB()
	defer db.Close(context.Background())
	user.Password = lib.Encrypt(user.Password)
	// fmt.Println(user)

	row := db.QueryRow(
		context.Background(),
		`insert into "users" (email, password) values ($1, $2) returning "id", "email", "password";
		insert into "profile" 'full_name' values($1)`,
		user.FullName, user.Email, user.Password,
	)

	var results FormRegister
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
	)
	return results
}