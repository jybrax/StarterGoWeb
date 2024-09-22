package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["authenticated"] == nil || !sess.Values["authenticated"].(bool) {
			// Redirige l'utilisateur vers la page de connexion s'il n'est pas authentifié
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}
