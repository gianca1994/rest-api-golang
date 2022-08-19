package main

import (
	"rest-api-golang/internal/database"
	"rest-api-golang/internal/logs"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	_ = logs.InitLogger()
	client := database.NewSqlClient("postgres://postgres:root@localhost:5432/rest-api-golang")
}
