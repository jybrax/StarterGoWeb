package controllers

import (
	"fmt"
	"wst/models"
	"wst/services" // Import du service via le module "wst"
)

func GetWeatherAll() ([]models.WeatherModel, error) {
	weatherData, err := services.GetWeatherJson()
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}

	return weatherData, nil
}
