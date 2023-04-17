package app

import (
	"log"
	"os"

	parsers "github.com/ivan-savchuk/pet-etl/pkg/parsers"
	pgdb "github.com/ivan-savchuk/pet-etl/pkg/pgdb"
)

func PutCurrentWeather() {
	log.Println("Putting Weather!")
	// get data for Kyiv
	lat := 50.45
	lon := 30.52
	wth, err := parsers.ParseWeather(lat, lon, os.Getenv("OPEN_WEATHER_API_KEY"))
	if err != nil {
		log.Fatalf("Response from OpenWeatherMap was not obtained: %v", err)
		return
	}

	pgdb.TruncateTable("weather.weather")
	pgdb.InsertIntoWeather(
		&wth.Weather[0].Description,
		&wth.Main.Temp,
		&wth.Main.FeelsLike,
		&wth.Main.Pressure,
		&wth.Main.Humidity,
		&wth.Wind.Speed,
		&lat,
		&lon,
	)
	log.Println("Successfully updated current weather data!")
}

func PutCurrentAQI() {
	log.Println("Processing AQI from OpenWeatherMap!")
	// Get data for Kyiv
	lat := 50.45
	lon := 30.52
	aqi, err := parsers.ParseAQI(lat, lon, os.Getenv("OPEN_WEATHER_API_KEY"))
	if err != nil {
		log.Fatalf("Response from OpenWeatherMap was not obtained: %v", err)
		return
	}

	pgdb.TruncateTable("weather.aqi")
	pgdb.InsertIntoAQI(
		&aqi.List[0].Main.Aqi,
		GetAQIDescription(aqi.List[0].Main.Aqi),
		&aqi.List[0].Components.Pm2_5,
		&lat,
		&lon,
	)
	log.Println("Successfully updated current AQI data!")
}
