package pgdb

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetPGConnection() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
