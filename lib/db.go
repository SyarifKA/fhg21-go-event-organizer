package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(),
		// "postgresql://postgres:1@103.93.58.89:54325/event_organizer?sslmode=disable",
		"postgresql://postgres:1@101.255.3.67:5432/event_organizer?sslmode=disable",
	)

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
