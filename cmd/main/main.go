package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ivan-savchuk/pet-etl/pkg/parsers"
)

var AQIAPIKey string = "undefined"

func init() {
	if value, ok := os.LookupEnv("AQI_API_KEY"); ok {
		AQIAPIKey = value
	}
}

func main() {
	if AQIAPIKey == "undefined" {
		log.Fatalln("'AQI_API_KEY' was not defined as environment variable.")
		return
	}
	resp, err := parsers.ParseAQI("Kyiv", "Kyiv", "Ukraine", AQIAPIKey)
	if err != nil {
		log.Fatalf("Response from AQI Api was not obtained: %v", err)
		return
	}
	if resp.Status == "fail" {
		log.Fatalf("Response from AQI Api was not obtained, response status: %v\n", resp.Status)
		return
	}

	fmt.Printf("AQI Response status: %v\n", resp.Status)
	fmt.Printf("Current AQI level in Kyiv: %v\n", resp.Data.Current.Pollution.Aqius)
}
