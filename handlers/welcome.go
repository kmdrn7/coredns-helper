package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetWelcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"message": "Selamat datang :)",
		})
	}
}