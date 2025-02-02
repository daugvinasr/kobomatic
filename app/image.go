package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/daugvinasr/kobomatic/database/calibre"
	"github.com/labstack/echo/v4"
)

func (app *App) GetImage(c echo.Context) error {
	path, err := calibre.FindBookCover(app.cq, c.Param("bookUUID"))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.NoContent(http.StatusNotFound)
		}
		return err
	}

	file, err := os.ReadFile(fmt.Sprintf("%s/%s/cover.jpg", app.env.LibraryFolder, path))
	if err != nil {
		return err
	}

	return c.Blob(http.StatusOK, "image/jpeg", file)
}
