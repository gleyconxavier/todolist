package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Activitie db
type Activitie struct {
	title string `db:"title"`
}

// ConnectDb to get stuff
func ConnectDb() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "user=gleyconxavier password=supersecret123 dbname=dbTodo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
