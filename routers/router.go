package routers

import (
	"net/http"
	"wst/controllers"
	"wst/libs"
	"wst/middlewares"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

// Router function receives the Echo instance to define the routes
func Router(e *echo.Echo) {
	e.Use(session.Middleware(store)) // Utiliser le middleware de session

	e.GET("/", home)
	e.GET("/about", about)
	e.GET("/weather", middlewares.AuthMiddleware(weather))
	e.GET("/user", user)
	e.GET("/login", login)               // Page de connexion
	e.POST("/submit-login", submitLogin) // Soumission du formulaire de connexion
	e.GET("/logout", logout)             // Route de déconnexion
	e.POST("/submit-form", controllers.SubmitFormHandler)
}

func submitLogin(c echo.Context) error {
	// Authentifier l'utilisateur
	err := controllers.SubmitLoginHandler(c)
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la soumission du login: %v", err)
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"error": "Nom d'utilisateur ou mot de passe incorrect",
		})
	}
	// exemple de valeur que peux ajouter dans la session
	sessionData := map[string]interface{}{
		"userID":    1,
		"userName":  "JohnDoe",
		"roles":     []string{"admin", "editor"},
		"loggedIn":  true,
		"lastLogin": "2024-09-23",
	}

	// Creation de la session utilisateur
	libs.CreateSession(c, sessionData)

	// Récupérer et afficher les données stockées
	_, data, err := libs.GetSessionData(c)
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la récupération de la session: %v", err)
	} else {
		c.Echo().Logger.Infof("Données de session récupérées: %v", data)
	}

	// Rediriger vers la page d'accueil après connexion
	return c.Redirect(http.StatusSeeOther, "/")
}

func logout(c echo.Context) error {
	// Supprimer la session
	libs.DeleteSession(c)

	return c.Redirect(http.StatusSeeOther, "/")
}

// Login handler
func login(c echo.Context) error {

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"ContentTemplate": "login.html",
	})
}

// Home handler (renders index.html)
func home(c echo.Context) error {
	authenticated, data, err := libs.GetSessionData(c)
	if err != nil {
		// Gérer le cas où l'utilisateur n'est pas authentifié
		authenticated = false
		data = map[string]interface{}{}
	}

	// Ajouter des informations supplémentaires à la page
	data["title"] = "Bienvenue sur go starter webpack"
	data["message"] = "Tu retrouveras tout pour créer ton application web."
	data["authenticated"] = authenticated
	data["userName"] = data["userName"]

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"ContentTemplate": "index.html",
		"Data":            data,
	})
}

func about(c echo.Context) error {
	authenticated, data, err := libs.GetSessionData(c)
	if err != nil {
		// Gérer le cas où l'utilisateur n'est pas authentifié
		authenticated = false
		data = map[string]interface{}{}
	}
	data["title"] = "About Page"
	data["message"] = "This is the about page"
	data["authenticated"] = authenticated
	data["userName"] = data["userName"]

	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"ContentTemplate": "about.html",
		"Data":            data,
	})
}

func user(c echo.Context) error {
	authenticated, data, err := libs.GetSessionData(c)
	if err != nil {
		// Gérer le cas où l'utilisateur n'est pas authentifié
		authenticated = false
		data = map[string]interface{}{}
	}
	data["title"] = "User Page"
	data["message"] = "This is the user page"
	data["authenticated"] = authenticated
	data["userName"] = data["userName"]

	return c.Render(http.StatusOK, "user.html", map[string]interface{}{
		"ContentTemplate": "user.html",
		"Data":            data,
	})
}

func weather(c echo.Context) error {
	authenticated, data, err := libs.GetSessionData(c)
	if err != nil {
		// Gérer le cas où l'utilisateur n'est pas authentifié
		authenticated = false
		data = map[string]interface{}{}
	}
	weatherData, err := controllers.GetWeatherAll()
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}
	data["title"] = "Weather Page"
	data["message"] = "Voici les informations météorologiques :"
	data["authenticated"] = authenticated
	data["weather"] = weatherData
	data["userName"] = data["userName"]

	return c.Render(http.StatusOK, "weather.html", map[string]interface{}{
		"ContentTemplate": "weather.html",
		"Data":            data,
	})
}
