package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Weather struct {
	List []struct {
		Dt     int64  `json:"dt"`
		DtText string `json:"dt_txt"`
		Main   struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	} `json:"list"`
}

func main() {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?lat=44.34&lon=10.99&appid=935a49f591a04de564e0cd1630a7c434")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available.")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	my_list := weather.List

	for _, l := range my_list {
		date := time.Unix(l.Dt, 0)

		fmt.Printf("%s - %s: %.0fC\n", date.Format("15:04"), l.DtText, l.Main.Temp)

	}
}
