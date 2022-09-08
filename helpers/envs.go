package helpers

import "os"

func URL_DB() string {
	var urlDb = os.Getenv("DATABASE_URL")
	if urlDb == "" {
		urlDb = "postgres://postgres:postgres@localhost:5432/qap?sslmode=disable"
	}
	return urlDb
}
