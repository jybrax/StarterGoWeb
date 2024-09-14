package controllers

import (
	"fmt"
	modelsWeather "wst/models"
	services "wst/services" // Import du service via le module "wst"
)

func GetWeatherAll() ([]modelsWeather.WeatherModel, error) {
	weatherData, err := services.GetWeatherJson()
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}

	return weatherData, nil
}
