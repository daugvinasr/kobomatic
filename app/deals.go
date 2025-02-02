package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) GetDeals(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Deals": []string{},
	})
}
