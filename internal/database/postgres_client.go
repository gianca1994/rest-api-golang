package database

import (
	"database/sql"
	"fmt"
)

type PostgresClient struct {

}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("postgres", source)
	if err != nil {
		_ = fmt.Errorf("error opening database: %s", err.Error())
		panic("..")
	}
	return db
}