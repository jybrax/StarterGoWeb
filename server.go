package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"wst/middlewares"
	"wst/routers"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom HTML template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	c.Echo().Logger.Infof("Rendering template: %s", name)

	// Lookup for the main layout template
	tmpl := t.templates.Lookup("base.html")
	if tmpl == nil {
		c.Echo().Logger.Errorf("base template not found: base.html")
		return echo.NewHTTPError(http.StatusInternalServerError, "base template not found")
	}

	// Check if the content template exists
	contentTmpl := t.templates.Lookup(name)
	if contentTmpl == nil {
		c.Echo().Logger.Errorf("Content template not found: %s", name)
		return echo.NewHTTPError(http.StatusInternalServerError, "Content template not found")
	}

	// Log the template and data for debugging
	c.Echo().Logger.Infof("Data for rendering: %+v", data)

	// Create a buffer to capture the content template output
	var contentBuffer bytes.Buffer

	// Render the content template into the buffer
	if err := contentTmpl.Execute(&contentBuffer, data); err != nil {
		return err
	}

	// Execute the main layout template with the rendered content
	return tmpl.ExecuteTemplate(w, "base.html", map[string]interface{}{
		"Content": template.HTML(contentBuffer.String()), // Mark as HTML
	})
}
func main() {
	e := echo.New()
	e.Use(middlewares.LoggingMiddleware)
	e.Static("/", "public")

	// Ajoute le gestionnaire d'erreur personnalisé
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		// Gérer les erreurs 404
		if code == http.StatusNotFound {
			c.Render(http.StatusNotFound, "404.html", map[string]interface{}{
				"message": "La page que vous cherchez est introuvable.",
			})
			return
		}

		// Pour les autres erreurs, affiche un message générique
		c.String(code, "Une erreur est survenue.")
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/layouts/*.html")),
	}
	renderer.templates = template.Must(renderer.templates.ParseGlob("public/view/*.html"))

	e.Renderer = renderer

	// Routes
	routers.Router(e)

	// Démarrer le serveur sur le port 1323
	e.Logger.Fatal(e.Start(":1323"))
}
