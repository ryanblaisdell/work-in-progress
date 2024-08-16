package main

import (
	"web-app-build/handlers"
	"web-app-build/middlewares"
	"web-app-build/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middlewares.RegisterMiddlewares(e)

	e.Static("/static", "static")

	e.Renderer = templates.NewTemplate()

	e.GET("/", handlers.IndexHandler)
	e.POST("/home", handlers.HomeHandler)
	e.POST("/about", handlers.AboutHandler)
	e.POST("/resume", handlers.ResumeHandler)
	e.POST("/contact", handlers.ContactHandler)
	e.POST("/login", handlers.LoginHandler)
	e.GET("/modal-content", handlers.ModalContentHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
