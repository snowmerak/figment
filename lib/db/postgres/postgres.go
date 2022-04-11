package postgres

import (
	"figment/lib/db"
	"os"

	"github.com/jackc/pgx"
)

var pool *pgx.ConnPool

func init() {
	pgxPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     5432,
			Database: os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		MaxConnections: 10,
	})
	if err != nil {
		panic(err)
	}

	pool = pgxPool
}

var _ db.Query = Query

func Query(sql string, args ...interface{}) ([]interface{}, error) {
	rows, err := pool.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	values, err := rows.Values()
	if err != nil {
		return nil, err
	}
	return values, nil
}
