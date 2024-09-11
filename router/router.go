package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Router function receives the Echo instance to define the routes
func Router(e *echo.Echo) {
	e.GET("/", home)
	e.GET("/about", about)
}

// Home handler (renders index.html)
func home(c echo.Context) error {
	data := map[string]interface{}{
		"title":   "Home Page",
		"message": "This is the home page",
	}
	c.Echo().Logger.Infof("Rendering / with data: %+v", data)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"ContentTemplate": "index.html",
		"Data":            data,
	})
}

// Login handler (renders login.html)
func about(c echo.Context) error {
	data := map[string]interface{}{
		"title":   "Login Page",
		"message": "This is the login page",
	}
	c.Echo().Logger.Infof("Rendering /about with data: %+v", data)
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"ContentTemplate": "about.html",
		"Data":            data,
	})
}
