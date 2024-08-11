package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Templates represents a collection of templates.
type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

// main is the entry point of the application.
// It initializes an instance of the Echo framework and sets up the necessary middleware.
// It also defines a route for the root URL ("/") that increments a count and renders a template.

func main() {
	// Create a new instance of the Echo framework.
	e := echo.New()

	// Use the Logger middleware to log HTTP requests and responses.
	e.Use(middleware.Logger())

	// Initialize a Count struct with a count value of 0.
	count := Count{Count: 0}

	e.Static("/static", "static")

	// Set the Renderer field of the Echo instance to a new template renderer.
	e.Renderer = newTemplate()

	// Define a route for the root URL ("/").
	e.GET("/", func(c echo.Context) error {
		// Render the "count" template with the updated count value.
		return c.Render(200, "index", count)
	})

	// Define a route for the root URL ("/home").
	e.POST("/home", func(c echo.Context) error {
		// Render the "count" template with the updated count value.
		return c.Render(200, "home", nil)
	})

	// Define a route for the root URL ("/about").
	e.POST("/about", func(c echo.Context) error {
		// Render the "count" template with the updated count value.
		return c.Render(200, "about", nil)
	})

	e.POST("/resume", func(c echo.Context) error {
		// Render the "count" template with the updated count value.
		return c.Render(200, "resume", nil)
	})

	e.POST("/contact", func(c echo.Context) error {
		// Render the "count" template with the updated count value.
		return c.Render(200, "contact", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
