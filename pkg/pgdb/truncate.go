package pgdb

import (
	"context"
	"log"
)

func TruncateTable(tableName string) {
	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "TRUNCATE TABLE " + tableName + ";"

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\nTruncated table %s successfully!", tableName)
}
