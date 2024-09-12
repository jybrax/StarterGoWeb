package controllersWeather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	modelsWeather "wst/models" // Corriger l'import pour utiliser modelsWeather
)

// Struct pour correspondre aux données JSON
type WeatherData struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
}

// GetWeather lit les données météo depuis un fichier JSON
func GetWeather() (modelsWeather.WeatherModel, error) { // Utiliser modelsWeather ici
	// Ouvrir le fichier JSON
	file, err := os.Open("data/weather.json")
	if err != nil {
		log.Printf("Erreur lors de l'ouverture du fichier JSON: %v", err)
		return modelsWeather.WeatherModel{}, err // Utiliser modelsWeather ici
	}
	defer file.Close()

	// Lire le contenu du fichier
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Erreur lors de la lecture du fichier JSON: %v", err)
		return modelsWeather.WeatherModel{}, err // Utiliser modelsWeather ici
	}

	// Parse JSON vers la struct WeatherData
	var weatherData WeatherData
	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		log.Printf("Erreur lors du parsing du JSON: %v", err)
		return modelsWeather.WeatherModel{}, err // Utiliser modelsWeather ici
	}

	// Retourner les données sous forme de modèle WeatherModel
	return modelsWeather.WeatherModel{
		Temperature: weatherData.Temperature,
		City:        weatherData.City,
		Weather:     weatherData.Weather,
	}, nil
}
