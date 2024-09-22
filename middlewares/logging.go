package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

// LoggingMiddleware ajoute un log pour chaque requête HTTP
func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		// Continue avec le handler suivant
		err := next(c)
		stop := time.Now()

		// Log des informations de la requête
		c.Logger().Infof("Méthode: %s, Chemin: %s, Temps d'exécution: %v",
			c.Request().Method, c.Request().URL.Path, stop.Sub(start))

		return err
	}
}
