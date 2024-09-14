package controllers

import (
	"net/http"
	"wst/models"
	"wst/services"

	"github.com/labstack/echo/v4"
)

func SubmitFormHandler(c echo.Context) error {
	var user models.UserModel
	user.Name = c.FormValue("name")
	user.UserName = c.FormValue("username")

	err := services.AddUserSql(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erreur lors de l'ajout de l'utilisateur")
	}

	return c.Redirect(http.StatusSeeOther, "/about?success=Le+formulaire+a+été+soumis+avec+succès!")
}
