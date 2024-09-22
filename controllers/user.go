package controllers

import (
	"net/http"
	"wst/models"
	"wst/services"

	"github.com/labstack/echo/v4"
)

func SubmitFormHandler(c echo.Context) error {
	var user models.UserModel
	user.UserName = c.FormValue("username")
	user.Password = c.FormValue("password")

	err := services.AddUserSql(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erreur lors de l'ajout de l'utilisateur")
	}

	return c.Redirect(http.StatusSeeOther, "/user?success=Le+formulaire+a+été+soumis+avec+succès!")
}
