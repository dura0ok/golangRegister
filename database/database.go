package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123321@/learn")

	if err != nil {
		return nil, err
	}
	return db, nil
}
