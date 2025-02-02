package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) PostAnalyticsTests(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Result":  "Success",
		"TestKey": "00000000-0000-0000-0000-000000000000",
		"Tests":   echo.Map{},
	})
}
