package app

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pgdb "github.com/ivan-savchuk/pet-etl/pkg/pgdb"
)

func GreetUser(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Hi, there! I'm weather bot! I will bring you some weather🌦️ and air polution data!" +
		"\n\nFeel free to use /help and discover what functions are available."

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func HelpUser(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "I understand /set and /current." +
		"\n\n🗺️ With /set you can specify your location." +
		"\n👐 With /current you can find out current weather and air polution." +
		"\n\nCurrent version of bot can only provide data for Kyiv, Ukraine 🇺🇦. Stay tuned!"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func SetUserData(bot *tgbotapi.BotAPI, update *tgbotapi.Update, updates *tgbotapi.UpdatesChannel) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	oldCityName := pgdb.SelectCityName(update.Message.Chat.ID)
	if !IsStringEmpty(oldCityName) {
		msg.Text = "You've already set your city to " + oldCityName + ". Your city name will be updated."
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	msg.Text = "Ok, now provide your city name.\n\nPs. I'm waiting 🧐"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	chatID := update.Message.Chat.ID
	var city string
	for {
		*update = <-*updates
		if update.Message != nil && update.Message.Chat.ID == chatID {
			city = update.Message.Text
			break
		}
	}

	cityID, cityName := pgdb.SelectClosestCity(&city)
	log.Println(cityID, cityName)
	if !IsStringEmpty(oldCityName) {
		pgdb.UpdateChatCity(chatID, int64(cityID))
	} else {
		pgdb.InsertIntoChats(chatID, int64(cityID))
	}

	msg.Text = fmt.Sprintf("Your curent location was set to **%s**", cityName)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func GetCurrentData(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	city1, aqi, pm2_5, descr1 := pgdb.SelectAQIData(update.Message.Chat.ID)
	if IsStringEmpty(city1) {
		msg.Text = "You did not set your city! Use /set to set your city, and then comeback here!"
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		return
	}
	city, descr, temp, feels_like, pressure, humidity, wind := pgdb.SelectWeatherData(update.Message.Chat.ID)

	newPressure := float64(pressure) * 0.75006
	fullMessage := fmt.Sprintf("Current weather 🌦️ in %s.\nDescription: %s\nTemparature: %2.f\nFeels like: %2.f\nPressure: %2.f\nHumidity: %d\nWind: %2.f",
		city, descr, temp, feels_like, newPressure, humidity, wind) +
		fmt.Sprintf("\n\nCurrent AQI rate 📊 in %s: %d\nDescription: %s\nPm2.5: %2.f", city1, aqi, descr1, pm2_5)
	msg.Text = fullMessage
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func DefaultResponse(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "I don't know that command"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
