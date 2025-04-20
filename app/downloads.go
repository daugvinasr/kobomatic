package app

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/daugvinasr/kobomatic/database/calibre"
	"github.com/labstack/echo/v4"
	"github.com/pgaskin/kepubify/v4/kepub"
)

func (app *App) GetBookFile(c echo.Context) error {
	row, err := calibre.FindBookFilePath(app.cq, c.Param("bookUUID"))
	if err != nil {
		return err
	}

	inputPath := fmt.Sprintf("%s/%s/%s.%s", app.env.LibraryFolder, row.Path, row.Name, "epub")

	output := bytes.NewBuffer(nil)

	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return nil
	}
	defer r.Close()

	err = kepub.NewConverter().Convert(context.Background(), output, r)
	if err != nil {
		return err
	}

	c.Response().Header().Set(
		"Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%s.kepub.epub\"",
			row.Name),
	)

	return c.Blob(http.StatusOK,
		"application/kepub",
		output.Bytes(),
	)
}
