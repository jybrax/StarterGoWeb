package test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wst/controllers"
	"wst/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// MockRenderer est un mock de renderer pour Echo
type MockRenderer struct{}

// Render est la fonction qui sera appelée pour rendre un template
func (m *MockRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Tu peux personnaliser ce qui est retourné ici
	_, err := w.Write([]byte("Rendered template"))
	return err
}

// Mock de la fonction VerifUserSql pour éviter d'interagir avec la vraie base de données
func mockVerifUserSql(user models.UserModel) error {
	if user.UserName == "validuser" && user.Password == "validpassword" {
		return nil
	}
	return errors.New("invalid username or password")
}

// TestSubmitLoginSuccess vérifie si le login réussit avec des informations correctes
func TestSubmitLoginSuccess(t *testing.T) {
	// Créer une instance d'Echo
	e := echo.New()

	// Assigner le renderer mock à Echo
	e.Renderer = &MockRenderer{}

	// Simuler une requête POST avec des données de formulaire valides
	req := httptest.NewRequest(http.MethodPost, "/submit-login", strings.NewReader("username=validuser&password=validpassword"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	// Simuler la réponse
	rec := httptest.NewRecorder()

	// Créer un contexte Echo pour la requête
	c := e.NewContext(req, rec)

	// Appeler la fonction `SubmitLogin` avec le mock
	err := controllers.SubmitLoginHandler(c, mockVerifUserSql)
	assert.NoError(t, err)

	// Vérifier que le code de réponse est bien 303 See Other
	assert.Equal(t, http.StatusOK, rec.Code)

}

// TestSubmitLoginFailure vérifie si le login échoue avec des informations incorrectes
func TestSubmitLoginFailure(t *testing.T) {
	// Créer une instance d'Echo
	e := echo.New()

	// Assigner le renderer mock à Echo
	e.Renderer = &MockRenderer{}

	// Simuler une requête POST avec des données de formulaire incorrectes
	req := httptest.NewRequest(http.MethodPost, "/submit-login", strings.NewReader("username=invaliduser&password=invalidpassword"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	// Simuler la réponse
	rec := httptest.NewRecorder()

	// Créer un contexte Echo pour la requête
	c := e.NewContext(req, rec)

	// Appeler la fonction `SubmitLogin` avec le mock
	err := controllers.SubmitLoginHandler(c, mockVerifUserSql)
	assert.NoError(t, err)

	// Vérifier que le code de réponse est bien 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

}
