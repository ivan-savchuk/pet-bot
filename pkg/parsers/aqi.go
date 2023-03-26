package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Country  string   `json:"country"`
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Current struct {
	Pollution Pollution `json:"pollution"`
	Weather   Weather   `json:"weather"`
}

type Pollution struct {
	Ts     string `json:"ts"`
	Aqius  int    `json:"aqius"`
	Mainus string `json:"mainus"`
	Aqicn  int    `json:"aqicn"`
	Maincn string `json:"maincn"`
}

type Weather struct {
	Ts string  `json:"ts"`
	Tp float64 `json:"tp"`
	Pr float64 `json:"pr"`
	Hu float64 `json:"hu"`
	Ws float64 `json:"ws"`
	Wd float64 `json:"wd"`
	Ic string  `json:"ic"`
}

func ParseAQI(city string, state string, country string, key string) (*Response, error) {

	response := &Response{}

	resp, err := http.Get(
		fmt.Sprintf(
			"http://api.airvisual.com/v2/city?city=%v&state=%v&country=%v&key=%v",
			city,
			state,
			country,
			key,
		),
	)
	if err != nil {
		log.Panicln(err)
		return response, err
	}

	defer resp.Body.Close()

	jsonBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
		return response, err
	}

	err = json.Unmarshal([]byte(jsonBody), response)
	if err != nil {
		log.Panicln(err)
		return response, err
	}
	return response, nil
}
