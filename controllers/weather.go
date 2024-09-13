package controllersWeather

import (
	"net/http"
	servicesWeather "wst/services" // Import du service via le module "wst"

	"github.com/labstack/echo/v4"
)

func GetWeatherHandler(c echo.Context) error {
	weatherData, err := servicesWeather.GetWeather("paris") // Appel au service
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Erreur lors de la récupération des données météo")
	}

	return c.Render(http.StatusOK, "weather.html", map[string]interface{}{
		"weather": weatherData,
	})
}
