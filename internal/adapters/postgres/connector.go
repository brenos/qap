package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/brenos/qap/helpers"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

func GetDbConnection() *sql.DB {
	db, err := sql.Open("postgres", helpers.URL_DB())
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 10)
	return db
}

// RunMigrations run scripts on path database/migrations
func RunMigrations() {
	databaseURL := helpers.URL_DB()
	m, err := migrate.New("file://database/migrations", "pgx"+databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
