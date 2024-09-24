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
	data := make(map[string]interface{})

	user := models.UserModel{
		UserName: username,
		Password: password,
	}

	err := verifUser(user)
	if err != nil {
		data["messageError"] = "Nom d'utilisateur ou mot de passe incorrect"
		log.Printf("Erreur d'authentification : %v", err)
		c.Echo().Logger.Infof("Rendering login.html with error message.")
		return c.Render(http.StatusUnauthorized, "login.html", map[string]interface{}{
			"ContentTemplate": "login.html",
			"Data":            data,
		})
	}

	return nil
}
