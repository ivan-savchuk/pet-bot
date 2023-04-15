package main

import (
	"fmt"
	"log"
	"os"
	"time"

	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ivan-savchuk/pet-etl/pkg/parsers"
	"github.com/ivan-savchuk/pet-etl/pkg/pgdb"
)

var WeatherAPIKey string = "undefined"
var TelegramToken string = "undefined"

func init() {
	if value, ok := os.LookupEnv("OPEN_WEATHER_API_KEY"); ok {
		WeatherAPIKey = value
	}

	if value, ok := os.LookupEnv("TELEGRAM_TOKEN"); ok {
		TelegramToken = value
	}
}

func main() {
	db, err := pgdb.GetPGConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var now time.Time

	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		panic(err)
	}

	fmt.Println(now)
	// bot, err := tgbotapi.NewBotAPI(TelegramToken)
	// if err != nil {
	// 	panic(err)
	// }

	// bot.Debug = true

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message != nil { // If we got a message
	// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		bot.Send(msg)
	// 	}
	// }
}

func getWeather() {
	if WeatherAPIKey == "undefined" {
		log.Fatalln("'OPEN_WEATHER_API_KEY' was not defined as environment variable.")
		return
	}
	// Get data for Kyiv
	wth, err := parsers.ParseWeather(50.45, 30.52, &WeatherAPIKey)
	if err != nil {
		log.Fatalf("Response from OpenWeatherMap was not obtained: %v", err)
		return
	}
	if wth.Cod != 200 {
		log.Fatalf("Response from OpenWeatherMap was not obtained, response status: %v\n", wth.Cod)
		return
	}
	// Get data for Kyiv
	aqi, err := parsers.ParseAQI(50.45, 30.52, &WeatherAPIKey)
	if err != nil {
		log.Fatalf("Response from OpenWeatherMap was not obtained: %v", err)
		return
	}

	fmt.Printf("OpenWeatherMap Response status: %v\n", wth.Cod)
	fmt.Printf("Current Weather in Kyiv: %v\n", wth.Weather[0].Description)
	fmt.Printf("Current temparature: %v\n", fmt.Sprintf("%.1f", wth.Main.Temp))
	fmt.Printf("Feels like: %v\n", fmt.Sprintf("%.1f", wth.Main.FeelsLike))
	fmt.Printf("Humidity: %v\n", wth.Main.Humidity)
	fmt.Printf("Pressure: %v\n", wth.Main.Pressure)
	fmt.Printf("Air pollution in Kyiv (AQI): %v\n", aqi.List[0].Main.Aqi)
	fmt.Println("Where 1 = Good, 2 = Fair, 3 = Moderate, 4 = Poor, 5 = Very Poor.")
}
