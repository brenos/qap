package helpers

import "os"

func URL_DB() string {
	var urlDb = os.Getenv("DATABASE_URL")
	if urlDb == "" {
		urlDb = "postgres://postgres:postgres@localhost:5432/qap?sslmode=disable"
	}
	return urlDb
}

func TOKEN_KEY() string {
	var key = os.Getenv("TOKEN_KEY")
	if key == "" {
		key = "1234567890123456"
	}
	return key
}

func TOKEN_EMAIL() string {
	var key = os.Getenv("TOKEN_EMAIL")
	return key
}
