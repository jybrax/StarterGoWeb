package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"wst/router"

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
	e.Static("/css", "css")
	e.Static("/js", "js")

	// Add logging middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e.Logger.Infof("Handling %s %s", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	})

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("layouts/*.html")),
	}
	renderer.templates = template.Must(renderer.templates.ParseGlob("view/*.html"))

	for _, tmpl := range renderer.templates.Templates() {
		e.Logger.Infof("Loaded template: %s", tmpl.Name())
	}

	e.Renderer = renderer

	// Set up routes
	router.Router(e)

	// Start the server on port 1323
	//e.Logger.SetLevel(log.DEBUG)

	e.Logger.Fatal(e.Start(":1323"))
}
