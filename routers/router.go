package routers

import (
	"net/http"
	"wst/controllers"

	"github.com/labstack/echo/v4"
)

// Router function receives the Echo instance to define the routes
func Router(e *echo.Echo) {
	e.GET("/", home)
	e.GET("/about", about)
	e.GET("/weather", weather)
	e.GET("/user", user)
	e.POST("/submit-form", controllers.SubmitFormHandler)
}

// Home handler (renders index.html)
func home(c echo.Context) error {
	data := map[string]interface{}{
		"title":   "Bienvenu sur go starter webpack",
		"message": "tu retrouvera tout pour créer ton application web. De plus tu peux utiliser la doc pour mieux comprendre sur /////",
	}
	c.Echo().Logger.Infof("Rendering / with data: %+v", data)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"ContentTemplate": "index.html",
		"Data":            data,
	})
}

// about handler (renders about.html)
func about(c echo.Context) error {
	data := map[string]interface{}{
		"title":   "About Page",
		"message": "This is the about page",
	}

	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"ContentTemplate": "about.html",
		"Data":            data,
	})
}

func user(c echo.Context) error {
	// Récupérer le message de succès s'il existe
	successMessage := c.QueryParam("success")

	data := map[string]interface{}{
		"title":   "User Page",
		"message": "This is the user page",
	}

	// Ajouter le message de succès dans les données, s'il existe
	if successMessage != "" {
		data["successMessage"] = successMessage
	}

	return c.Render(http.StatusOK, "user.html", map[string]interface{}{
		"ContentTemplate": "user.html",
		"Data":            data,
	})
}

func weather(c echo.Context) error {
	weatherData, err := controllers.GetWeatherAll()

	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}

	data := map[string]interface{}{
		"title":   "Weather Page",
		"message": "Voici les informations météorologiques :",
		"weather": weatherData, // Assurez-vous que weatherData est une slice
	}
	c.Echo().Logger.Infof("Rendering /weather avec data: %+v", data)
	return c.Render(http.StatusOK, "weather.html", map[string]interface{}{
		"ContentTemplate": "weather.html",
		"Data":            data,
	})
}
