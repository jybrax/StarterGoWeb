package servicesWeather

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	database "wst/libs"
	modelsWeather "wst/models"

	"github.com/joho/godotenv"
)

// GetWeather récupère les données météo pour une ville
func GetWeather() (modelsWeather.WeatherModel, error) {
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

	// Requête SQL pour récupérer la météo d'une ville
	query := `SELECT * FROM Meteo`

	// Créer une instance vide du modèle météo
	var weather modelsWeather.WeatherModel

	// Utiliser la connexion à la base de données depuis le package config
	err = db.QueryRow(query).Scan(&weather.City, &weather.Temperature, &weather.Weather, &weather.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return weather, fmt.Errorf("Aucune donnée trouvée")
		}
		return weather, fmt.Errorf("Erreur lors de la récupération de la météo: %v", err)
	}

	return weather, nil
}
