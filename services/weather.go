package servicesWeather

import (
	"encoding/json"
	"io/ioutil"
	"os"
	modelsWeather "wst/models" // Import correct du modèle via le module
)

func GetWeather() (modelsWeather.WeatherModel, error) {
	file, err := os.Open("data/weather.json")
	if err != nil {
		return modelsWeather.WeatherModel{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return modelsWeather.WeatherModel{}, err
	}

	var weather modelsWeather.WeatherModel
	if err := json.Unmarshal(data, &weather); err != nil {
		return modelsWeather.WeatherModel{}, err
	}

	return weather, nil
}
