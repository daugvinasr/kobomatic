package app

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/daugvinasr/kobomatic/database/calibre"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (app *App) GetBookFile(c echo.Context) error {
	row, err := calibre.FindBookFilePath(app.cq, c.Param("bookUUID"))
	if err != nil {
		return err
	}

	inputPath := fmt.Sprintf("%s/%s/%s.%s", app.env.LibraryFolder, row.Path, row.Name, "epub")

	_, err = os.Stat(inputPath)
	if err != nil {
		return err
	}

	outputPath := fmt.Sprintf("/tmp/%s.kepub.epub", uuid.New())

	err = exec.Command("kepubify", inputPath, "-o", outputPath).Run()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(outputPath)
	if err != nil {
		return err
	}

	defer os.Remove(outputPath)

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.kepub.epub\"", row.Name))

	return c.Blob(http.StatusOK, "application/kepub", file)
}
