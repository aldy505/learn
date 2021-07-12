package main

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type KV struct {
	key   string `db:"key"`
	value string `db:"value"`
}

func main() {
	psq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	poolConfig, err := pgxpool.ParseConfig("postgres://postgres:password@localhost:5432/learn_go")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	query := `CREATE TABLE IF NOT EXISTS test (
		key varchar(255) UNIQUE, 
		values varchar(255)
	);`
	_, err = conn.Query(context.Background(), query)
	if err != nil {
		log.Fatalln(err)
	}

	query, args, err := psq.Insert("test").Columns("key", "values").Values("hello", "world").Values("intelligence", "fbi").ToSql()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = conn.Query(context.Background(), query, args...)
	if err != nil {
		log.Fatalln(err)
	}

	query, args, err = psq.Select("*").From("test").ToSql()
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := conn.Query(context.Background(), query, args...)
	if err != nil {
		log.Fatalln(err)
	}
}
