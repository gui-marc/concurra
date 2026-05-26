package repository

import (
	"context"
	"fmt"

	internal "github.com/gui-marc/concurra/repository/internal/generated/pgx"
	"github.com/jackc/pgx/v5"
)

func NewPGQueries(databaseURL string) (*internal.Queries, *pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	return internal.New(conn), conn, nil
}
