package routers

import (
	"net/http"
	"wst/controllers"
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
	sess, _ := session.Get("session", c)

	// Authentifier l'utilisateur
	err := controllers.SubmitLoginHandler(c)
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la soumission du login: %v", err)
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"error": "Nom d'utilisateur ou mot de passe incorrect",
		})
	}

	// Enregistrer l'utilisateur comme authentifié dans la session
	sess.Values["authenticated"] = true // Assurez-vous que cette ligne est bien là
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la sauvegarde de la session: %v", err)
	}

	// Rediriger vers la page d'accueil après connexion
	return c.Redirect(http.StatusSeeOther, "/")
}

func logout(c echo.Context) error {
	// Supprimer la session
	session, _ := store.Get(c.Request(), "session")
	session.Options.MaxAge = -1 // Expirer immédiatement la session
	session.Save(c.Request(), c.Response())

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
	sess, _ := session.Get("session", c)
	authenticated := sess.Values["authenticated"] != nil && sess.Values["authenticated"].(bool)

	data := map[string]interface{}{
		"title":         "Bienvenue sur go starter webpack",
		"message":       "Tu retrouveras tout pour créer ton application web.",
		"authenticated": authenticated,
	}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"ContentTemplate": "index.html",
		"Data":            data,
	})
}

func about(c echo.Context) error {
	sess, _ := session.Get("session", c)
	authenticated := sess.Values["authenticated"] != nil && sess.Values["authenticated"].(bool)

	data := map[string]interface{}{
		"title":         "About Page",
		"message":       "This is the about page",
		"authenticated": authenticated,
	}

	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"ContentTemplate": "about.html",
		"Data":            data,
	})
}

func user(c echo.Context) error {
	sess, _ := session.Get("session", c)
	authenticated := sess.Values["authenticated"] != nil && sess.Values["authenticated"].(bool)

	data := map[string]interface{}{
		"title":         "User Page",
		"message":       "This is the user page",
		"authenticated": authenticated,
	}

	return c.Render(http.StatusOK, "user.html", map[string]interface{}{
		"ContentTemplate": "user.html",
		"Data":            data,
	})
}

func weather(c echo.Context) error {
	sess, _ := session.Get("session", c)
	authenticated := sess.Values["authenticated"] != nil && sess.Values["authenticated"].(bool)

	weatherData, err := controllers.GetWeatherAll()
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la récupération des données météo: %v", err)
	}

	data := map[string]interface{}{
		"title":         "Weather Page",
		"message":       "Voici les informations météorologiques :",
		"weather":       weatherData,
		"authenticated": authenticated,
	}
	return c.Render(http.StatusOK, "weather.html", map[string]interface{}{
		"ContentTemplate": "weather.html",
		"Data":            data,
	})
}
