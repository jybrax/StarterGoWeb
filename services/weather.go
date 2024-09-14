package servicesWeather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	database "wst/libs"
	modelsWeather "wst/models"

	"github.com/joho/godotenv"
)

// GetWeather récupère toutes les données météo
func GetWeatherSql() ([]modelsWeather.WeatherModel, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	sqlData := database.SqlData{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		DataBaseName: os.Getenv("DB_NAME"),
	}

	db, err := sqlData.ConnectMysql()
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données:", err)
	}
	defer db.Close()

	// Requête SQL pour récupérer toutes les lignes de la table Meteo
	query := `SELECT * FROM Meteo`

	// Créer une slice pour stocker les résultats
	var weatherData []modelsWeather.WeatherModel

	// Exécuter la requête
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}
	defer rows.Close()

	// Parcourir les résultats
	for rows.Next() {
		var weather modelsWeather.WeatherModel
		if err := rows.Scan(&weather.City, &weather.Temperature, &weather.Weather, &weather.Date); err != nil {
			return nil, fmt.Errorf("Erreur lors de la récupération des données météo: %v", err)
		}
		weatherData = append(weatherData, weather)
	}

	// Vérifier les erreurs de parcours
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}
	return weatherData, nil
}

func GetWeatherJson() ([]modelsWeather.WeatherModel, error) {
	// Lire le fichier JSON
	file, err := ioutil.ReadFile("data/weather.json")
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la lecture du fichier JSON: %v", err)
	}

	// Créer une slice pour stocker les résultats
	var weatherData []modelsWeather.WeatherModel

	// Convertir le JSON en slice de WeatherModel
	if err := json.Unmarshal(file, &weatherData); err != nil {
		return nil, fmt.Errorf("Erreur lors de la conversion du JSON: %v", err)
	}

	return weatherData, nil
}
