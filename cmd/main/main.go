package main

import (
	"log"
	"os"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	app "github.com/ivan-savchuk/pet-etl/pkg/app"
	scheduler "github.com/ivan-savchuk/pet-etl/pkg/scheduler"
)

var WeatherAPIKey string = "undefined"
var TelegramToken string = "undefined"

func init() {
	if value, ok := os.LookupEnv("OPEN_WEATHER_API_KEY"); ok {
		WeatherAPIKey = value
	}
	if WeatherAPIKey == "undefined" {
		log.Panic("'OPEN_WEATHER_API_KEY' was not defined as environment variable.")
	}

	if value, ok := os.LookupEnv("TELEGRAM_TOKEN"); ok {
		TelegramToken = value
	}

	if TelegramToken == "undefined" {
		log.Panic("'OPEN_WEATHER_API_KEY' was not defined as environment variable.")
	}
}

func main() {
	location, err := time.LoadLocation("Europe/Kiev")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	sched := scheduler.Scheduler{}
	sched.SetNewScheduler(location)
	sched.AddNewJob("0 */8 * * *", app.PutCurrentAQI)
	sched.AddNewJob("0 */8 * * *", app.PutCurrentWeather)
	sched.Start()
	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		// user flow located here
		switch update.Message.Command() {
		case "start":
			go app.GreetUser(bot, update)
		case "help":
			go app.HelpUser(bot, update)
		case "set":
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				app.SetUserData(bot, &update, &updates)
			}()
			wg.Wait()
		case "current":
			go app.GetCurrentData(bot, update)
		default:
			go app.DefaultResponse(bot, update)
		}
	}
}
