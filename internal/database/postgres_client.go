package database

import (
	"database/sql"
)

type PostgresClient struct {
	*sql.DB
}

func NewSqlClient(source string) *PostgresClient {
	db, err := sql.Open("pgx", source)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresClient{db}
}
