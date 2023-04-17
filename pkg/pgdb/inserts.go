package pgdb

import (
	"context"
	"fmt"
	"log"
	"strconv"
)

func InsertIntoChats(chatID int64, cityID int64) {
	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	strChatID := strconv.FormatInt(chatID, 10)
	strCityID := strconv.FormatInt(cityID, 10)
	query := "INSERT INTO weather.chats (chat_id, city_id, last_activity)" +
		"VALUES ('" + strChatID + "', " + strCityID + ", NOW())"

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Println("Can't execute insert!")
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Chat data inserted successfully!")
}

func InsertIntoAQI(aqi *int, descr string, pm25 *float64, lat *float64, lon *float64) {
	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	strAQI := fmt.Sprintf("%d", *aqi)
	strPM := fmt.Sprintf("%f", *pm25)
	pointLoc := fmt.Sprintf("POINT(%f, %f)", *lat, *lon)
	query := "INSERT INTO weather.aqi (aqi, pm_25, descr, search_point, aqi_time)" +
		"VALUES (" + strAQI + ", " + strPM + ", '" + descr + "', " + pointLoc + ", NOW());"

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Println("Can't execute insert!")
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AQI data inserted successfully!")
}

func InsertIntoWeather(
	descr *string,
	temp *float64,
	feels_like *float64,
	pressure *int,
	humidity *int,
	wind *float64,
	lat *float64,
	lon *float64) {

	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	strTemp := fmt.Sprintf("%f", *temp)
	strFL := fmt.Sprintf("%f", *feels_like)
	strPressure := fmt.Sprintf("%d", *pressure)
	strHumidity := fmt.Sprintf("%d", *humidity)
	strWind := fmt.Sprintf("%f", *wind)
	pointLoc := fmt.Sprintf("POINT(%f, %f)", *lat, *lon)

	query := "INSERT INTO weather.weather (descr, temparature, feels_like, pressure, humidity, wind_speed, search_point, weather_time)" +
		"\nVALUES ('" + *descr + "', " + strTemp + ", " + strFL + ", " + strPressure + ", " + strHumidity + ", " + strWind + ", " + pointLoc + ", NOW());"

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		log.Println("Can't execute insert!")
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Weather data inserted successfully!")
}
