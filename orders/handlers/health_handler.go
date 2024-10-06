package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy!")
}
