package pgdb

import (
	"log"
	"strconv"
)

func SelectClosestCity(city *string) (int, string) {
	query := "SELECT city_id, city_name" +
		"\nFROM weather.cities" +
		"\nWHERE city_name % '" + *city + "'" +
		"\nORDER BY similarity(city_name, '" + *city + "') DESC" +
		"\nLIMIT 1;"

	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var cityID int
	var cityName string
	err = db.QueryRow(query).Scan(&cityID, &cityName)
	if err != nil {
		panic(err)
	}
	return cityID, cityName
}

func SelectCityName(chatID int64) string {
	strChatID := strconv.FormatInt(chatID, 10)
	query := "SELECT city_name" +
		"\nFROM weather.chats" +
		"\nINNER JOIN weather.cities" +
		"\nUSING (city_id)" +
		"\nWHERE chat_id = '" + strChatID + "';"

	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var cityName string
	err = db.QueryRow(query).Scan(&cityName)
	if err != nil {
		log.Println(err)
		return ""
	}
	return cityName
}

func SelectAQIData(chatID int64) (string, int, float64, string) {
	strChatID := strconv.FormatInt(chatID, 10)
	query :=
		"WITH loc AS (" +
			"\n\tSELECT city_name, city_location" +
			"\n\tFROM weather.chats" +
			"\n\tINNER JOIN weather.cities" +
			"\n\t\tUSING(city_id)" +
			"\n\tWHERE chat_id = '" + strChatID + "'" +
			"\n)" +
			"\nSELECT city_name, aqi, pm_25, descr" +
			"\nFROM weather.aqi, loc" +
			"\nORDER BY aqi.search_point <-> loc.city_location" +
			"\nLIMIT 1;"

	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var city string
	var aqi int
	var pm2_5 float64
	var descr string
	err = db.QueryRow(query).Scan(&city, &aqi, &pm2_5, &descr)
	if err != nil {
		log.Println(err)
		return "", 0, 0.0, ""
	}
	return city, aqi, pm2_5, descr
}

func SelectWeatherData(chatID int64) (string, string, float64, float64, int, int, float64) {
	strChatID := strconv.FormatInt(chatID, 10)
	query :=
		"WITH loc AS (" +
			"\n\tSELECT city_name, city_location" +
			"\n\tFROM weather.chats" +
			"\n\tINNER JOIN weather.cities" +
			"\n\t\tUSING(city_id)" +
			"\n\tWHERE chat_id = '" + strChatID + "'" +
			"\n)" +
			"\nSELECT city_name, descr, temparature, feels_like, pressure, humidity, wind_speed" +
			"\nFROM weather.weather, loc" +
			"\nORDER BY weather.search_point <-> loc.city_location" +
			"\nLIMIT 1;"

	db, err := GetPGConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var city string
	var descr string
	var temp float64
	var feels_like float64
	var pressure int
	var humidity int
	var wind float64
	err = db.QueryRow(query).Scan(&city, &descr, &temp, &feels_like, &pressure, &humidity, &wind)
	if err != nil {
		log.Println(err)
		return "", "", 0.0, 0.0, 0, 0, 0.0
	}
	return city, descr, temp, feels_like, pressure, humidity, wind
}
