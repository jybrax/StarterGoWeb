package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SubmitFormHandler(c echo.Context) error {
	// Récupérer les données du formulaire
	name := c.FormValue("name")
	username := c.FormValue("username")
	preference := c.FormValue("preference")
	remember := c.FormValue("remember")

	// Traiter les données comme tu le souhaites (enregistrer en base de données, etc.)
	fmt.Print(name + username + preference + remember)
	// Rediriger vers /about avec un message de succès
	return c.Redirect(http.StatusSeeOther, "/about?success=Le+formulaire+a+été+soumis+avec+succès!")
}
