package parsers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AQIData struct {
	List []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Aqi int `json:"aqi"`
		} `json:"main"`
		Components struct {
			Co    float64 `json:"co"`
			No    float64 `json:"no"`
			No2   float64 `json:"no2"`
			O3    float64 `json:"o3"`
			So2   float64 `json:"so2"`
			Pm2_5 float64 `json:"pm2_5"`
			Pm10  float64 `json:"pm10"`
			Nh3   float64 `json:"nh3"`
		} `json:"components"`
	} `json:"list"`
}

func ParseAQI(lat float32, lon float32, apiKey *string) (*AQIData, error) {
	response := &AQIData{}

	resp, err := http.Get(
		fmt.Sprintf(
			"https://api.openweathermap.org/data/2.5/air_pollution?lat=%v&lon=%v&APPID=%v",
			lat,
			lon,
			*apiKey,
		),
	)

	if err != nil {
		log.Panicln(err)
		return response, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		log.Panicln(err)
		return response, err
	}

	err = json.Unmarshal(buf.Bytes(), response)
	if err != nil {
		log.Panicln(err)
		return response, err
	}
	return response, nil
}
