package models

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func Connection() (*pgx.Conn, error) {
	databaseURL := "postgres://postgres:password@localhost:5432/learn_go"
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())
	return conn, nil
}

// func Setup() error {
// TODO: create a table if not exists, insert data from data/query.sql
// }
