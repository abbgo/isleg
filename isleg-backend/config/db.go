package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB *pgxpool.Pool

func ConnDB() (*pgxpool.Pool, error) {

	fmt.Println(os.Getenv("DATABASE_URL"))

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err := dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	DB = dbpool

	return dbpool, nil

}
