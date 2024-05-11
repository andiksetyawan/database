package main

import (
	"fmt"
	"log"

	"github.com/andiksetyawan/database"
	"github.com/andiksetyawan/database/sqlx"
)

func main() {
	db, err := sqlx.New(sqlx.WithPostgres(database.Config{
		Database: "foo",
		User:     "postgres",
		Host:     "localhost",
		Port:     "5432",
		Password: "secret",
	}))
	if err != nil {
		log.Fatal(fmt.Errorf("fail to connecting psql database : %w", err))
	}

	user := struct {
		ID   int
		Name string
	}{}

	_ = db.GetMaster().Get(&user, "SELECT * FROM user")
}
