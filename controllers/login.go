package controllers

import (
	"log"
	"net/http"
	"wst/models"
	"wst/services"

	"github.com/labstack/echo/v4"
)

func SubmitLoginHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.UserModel{
		UserName: username,
		Password: password,
	}

	err := services.VerifUserSql(user)
	if err != nil {
		log.Printf("Erreur d'authentification : %v", err)
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"error": "Nom d'utilisateur ou mot de passe incorrect",
		})
	}

	return nil
}
