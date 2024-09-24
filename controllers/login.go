package controllers

import (
	"log"
	"net/http"
	"wst/models"

	"github.com/labstack/echo/v4"
)

type VerifUserFunc func(user models.UserModel) error

func SubmitLoginHandler(c echo.Context, verifUser VerifUserFunc) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.UserModel{
		UserName: username,
		Password: password,
	}

	err := verifUser(user)
	if err != nil {
		log.Printf("Erreur d'authentification : %v", err)
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"error": "Nom d'utilisateur ou mot de passe incorrect",
		})
	}

	return nil
}
