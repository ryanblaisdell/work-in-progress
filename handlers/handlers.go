package handlers

import (
	"github.com/labstack/echo/v4"
)

type Count struct {
	Count int
}

func IndexHandler(c echo.Context) error {
	count := Count{Count: 0}
	return c.Render(200, "index", count)
}

func HomeHandler(c echo.Context) error {
	return c.Render(200, "home", nil)
}

func AboutHandler(c echo.Context) error {
	return c.Render(200, "about", nil)
}

func ResumeHandler(c echo.Context) error {
	return c.Render(200, "resume", nil)
}

func ContactHandler(c echo.Context) error {
	return c.Render(200, "contact", nil)
}

func LoginHandler(c echo.Context) error {
	return c.Render(200, "login", nil)
}

func ModalContentHandler(c echo.Context) error {
	return c.Render(200, `modal-content`, nil)
}
