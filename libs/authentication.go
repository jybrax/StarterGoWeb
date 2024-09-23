package libs

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func CreateSession(c echo.Context, sessionData map[string]interface{}) {
	sess, _ := session.Get("session", c)

	sess.Values["data"] = sessionData
	sess.Values["authenticated"] = true

	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		c.Echo().Logger.Errorf("Erreur lors de la sauvegarde de la session: %v", err)
	}
}

func GetSessionData(c echo.Context) (bool, map[string]interface{}, error) {
	sess, _ := session.Get("session", c)

	// Vérifier si l'utilisateur est authentifié
	authenticated, ok := sess.Values["authenticated"].(bool)
	if !ok || !authenticated {
		return false, nil, fmt.Errorf("Utilisateur non authentifié")
	}

	// Récupérer les données stockées dans la map
	if data, ok := sess.Values["data"].(map[string]interface{}); ok {
		return authenticated, data, nil
	}

	return authenticated, nil, fmt.Errorf("Aucune donnée trouvée dans la session")
}

func VeryfiAuth(c echo.Context) bool {
	// Récupérer la session via la fonction GetSessionData
	_, data, err := GetSessionData(c)

	if err != nil {
		// Si l'utilisateur n'est pas authentifié ou aucune donnée trouvée
		c.Echo().Logger.Errorf("Erreur de récupération des données de session : %v", err)
		return false
	}

	// Vérifier si l'utilisateur est authentifié
	if authenticated, ok := data["authenticated"].(bool); ok && authenticated {
		return true
	}

	return false
}

func DeleteSession(c echo.Context) {
	session, _ := store.Get(c.Request(), "session")
	session.Options.MaxAge = -1 // Expirer immédiatement la session
	session.Save(c.Request(), c.Response())
}
