package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		TimeEpoch int64  `json:"localtime_epoch"`
	} `json:"location"`
	Current struct {
		Tempc     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Feelslike  float32 `json:"feelslike_c"`
		Airquality struct {
			Pm10 float64 `json:"pm10"`
		} `json:"air_quality"`
	} `json:"current"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	api := os.Getenv("api")
	res, err := http.Get(api)
	if err != nil {
		panic(api)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("API not available")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("err")
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	location, current := weather.Location, weather.Current
	fmt.Printf("%s, %s:%.0fC, %s\nAQI:%.0f\n", location.Name, location.Country, current.Tempc, current.Condition.Text, current.Airquality.Pm10)
}
