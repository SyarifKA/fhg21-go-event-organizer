package models

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int     `json:"id"`
	Email    string  `json:"email" form:"email" binding:"required,email"`
	Password string  `json:"password" form:"password" binding:"required,min=8"`
	Username *string `json:"username" binding:"omitempty"`
}

type Password struct {
	// OldPassword     string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword" db:"password"`
	// ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

// var Data = []User{
// 	{
// 		Id:       1,
// 		Name:     "syarif",
// 		Email:    "syarif@mail.com",
// 		Password: "12345678",
// 	},
// }

// func FindAllUsers(dataSource []User) []User {
// 	return dataSource
// }

// func CreateUser(data User) User {
// 	id := 0
// 	for _, item := range Data {
// 		id = item.Id
// 	}
// 	data.Id = id + 1
// 	Data = append(Data, data)
// 	return data
// }

// func GetOneUser(dataSource []User, id int) User {
// 	index := -1
// 	for i, v := range dataSource {
// 		if v.Id == id {
// 			index = i
// 		}
// 	}
// 	return dataSource[index]
// }

// func UpdateDataById(data User, id int) User {
// 	ids := -1
// 	for index, item := range Data {
// 		if id == item.Id {
// 			ids = index
// 		}
// 	}
// 	if ids != 0 {
// 		Data[ids].Name = data.Name
// 		Data[ids].Email = data.Email
// 		Data[ids].Password = data.Password
// 		data.Id = Data[ids].Id
// 	}
// 	return data
// }

// func DeleteUser(id int) User {
// 	selected := -1
// 	userDelete := User{}
// 	for ids, item := range Data {
// 		if item.Id == id {
// 			selected = ids
// 			userDelete = item
// 		}
// 	}
// 	leftSide := Data[:selected]
// 	rightSide := Data[selected+1:]
// 	if userDelete.Id != 0 {
// 		Data = append(leftSide, rightSide...)
// 	}
// 	return userDelete
// }

// CRUD users

func FindAllUsers(search string, limit int, page int) ([]User, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	inputSQL := `select * from "users" where "email" ilike '%' || $1 || '%' limit $2 offset $3`
	rows, _ := db.Query(context.Background(), inputSQL, search, limit, offset)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])
	if err != nil {
		fmt.Println(err)
	}
	count := TotalData(search)
	return users, count
}

func TotalData(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())
	inputSQL := `select count(id) as "total" from "users" where "email" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), inputSQL, search)
	var result int
	rows.Scan(
		&result,
	)
	return result
}

func FindOneUserById(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "users" where "id" = $1`,
		id,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, item := range users {
		if item.Id == id {
			user = item
		}
	}
	return user
}

func FindOneUserByEmail(email string) User {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "users" where "email"=$1`,
		email,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, item := range users {
		if item.Email == email {
			user = item
		}
	}
	return user
}

func CreateUser(user User) User {
	db := lib.DB()
	defer db.Close(context.Background())
	user.Password = lib.Encrypt(user.Password)
	// fmt.Println(user)

	row := db.QueryRow(
		context.Background(),
		`insert into "users" (email, password, username) values ($1, $2, $3) returning "id", "email", "password", "username"`,
		user.Email, user.Password, user.Username,
	)

	var results User
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.Username,
	)
	return results
}

func DeleteUser(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag1, err1 := db.Exec(
		context.Background(),
		`delete from "users" where "id" = $1`,
		id,
	)
	commandTag2, err2 := db.Exec(
		context.Background(),
		`delete from "profile" where "id" = $1`,
		id,
	)

	if err1 != nil {
		return fmt.Errorf("failed to execute delete")
	}
	if err2 != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag1.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	if commandTag2.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func EditUser(user User, id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	user.Password = lib.Encrypt(user.Password)

	dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

	edit, err := db.Query(context.Background(), dataSql, user.Email, user.Username, user.Password, id)

	if err != nil {
		fmt.Println(err)
	}
	var result User
	edit.Scan()
	// user.Id = id
	return result
}

func EditPassword(user Password, id int) User {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(user.NewPassword)
	fmt.Println(id)
	user.NewPassword = lib.Encrypt(user.NewPassword)
	sql := `update "users" set "password" values $1 where "id" = $2 returning "id", "password"`

	change := db.QueryRow(context.Background(), sql, user.NewPassword, id)

	var result User

	change.Scan(
		&result.Id,
		&result.Email,
		&result.Password,
	)
	return result
}
