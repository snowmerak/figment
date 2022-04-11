package db

import (
	"figment/lib/dotenv"
)

func init() {
	if err := dotenv.Load("env/account_db.env"); err != nil {
		panic(err)
	}
}

type Query func(sql string, args ...interface{}) ([]interface{}, error)
