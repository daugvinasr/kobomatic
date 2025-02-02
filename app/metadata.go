package app

import (
	"net/http"

	"github.com/daugvinasr/kobomatic/database/calibre"
	"github.com/daugvinasr/kobomatic/gen"
	"github.com/labstack/echo/v4"
)

func (app *App) GetBookMetadata(c echo.Context) error {
	row, err := calibre.FindBookMetadata(app.cq, c.Param("bookUUID"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, []echo.Map{gen.Metadata(app.env.ServerAddress, row)})
}
