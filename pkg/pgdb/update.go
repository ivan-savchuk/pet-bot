package pgdb

import (
	"context"
	"log"
	"strconv"
)

func UpdateChatCity(chatID int64, cityID int64) {
	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	strChatID := strconv.FormatInt(chatID, 10)
	strCityID := strconv.FormatInt(cityID, 10)

	query := "UPDATE weather.chats " +
		"\nSET city_id = " + strCityID +
		"\nWHERE chat_id = '" + strChatID + "'"

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		if rb := tx.Rollback(); rb != nil {
			log.Fatalf("query failed: %v, unable to abort: %v", err, rb)
		}
		log.Fatal(err)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Chat data updated successfully!")
}
