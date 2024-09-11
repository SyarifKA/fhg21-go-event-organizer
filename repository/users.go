package repository

import (
	"context"
	"fmt"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllUsers(search string, limit int, page int) ([]dtos.User, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	inputSQL := `select * from "users" where "email" ilike '%' || $1 || '%' limit $2 offset $3`
	rows, _ := db.Query(context.Background(), inputSQL, search, limit, offset)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.User])
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

func FindOneUserById(id int) dtos.User {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "users" where "id" = $1`,
		id,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.User])

	if err != nil {
		fmt.Println(err)
	}

	user := dtos.User{}
	for _, item := range users {
		if item.Id == id {
			user = item
		}
	}
	return user
}

func FindOneUserByEmail(email string) dtos.User {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "users" where "email"=$1`,
		email,
	)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.User])

	if err != nil {
		fmt.Println(err)
	}

	user := dtos.User{}
	for _, item := range users {
		if item.Email == email {
			user = item
		}
	}
	return user
}

func CreateUser(data models.Users) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	sql := `
		INSERT INTO users (email, password)
		VALUES ($1, $2) RETURNING *
		`
	row, err := db.Query(context.Background(), sql, data.Email, data.Password)

	if err != nil {
		return models.Users{}, nil
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[models.Users])
	fmt.Println(user)
	if err != nil {
		return models.Users{}, nil
	}

	return user, err
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

func EditUser(user models.UserUpdate, id int) (*models.UserUpdate, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	// user.Password = lib.Encrypt(user.Password)

	dataSql := `update "users" set (email , username) = ($1, $2) where id=$3 returning "id", "email", "username"`

	row, err := db.Query(context.Background(), dataSql, user.Email, user.Username, id)
	if err != nil {
		fmt.Println(err)
	}
	result, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByPos[models.UserUpdate])
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

func EditPassword(user models.UpdatePassword, id int) (models.UserUpdate, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	fmt.Println(user.NewPassword)
	fmt.Println(id)
	user.NewPassword = lib.Encrypt(user.NewPassword)
	sql := `update users set password = $1 where id = $2 returning *`

	change, err := db.Query(context.Background(), sql, user.NewPassword, id)

	if err != nil {
		return models.UserUpdate{}, err
	}

	result, err := pgx.CollectOneRow(change, pgx.RowToStructByPos[models.UserUpdate])

	if err != nil {
		return models.UserUpdate{}, err
	}

	return result, err
}
