package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) GetCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"ItemCount":        10,
		"TotalPageCount":   1,
		"TotalItemCount":   10,
		"CurrentPageIndex": 0,
		"ItemsPerPage":     1000,
	})
}
