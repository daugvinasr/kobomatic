package app

import (
	"database/sql"
	"net/http"

	"github.com/daugvinasr/kobomatic/env"
	"github.com/labstack/echo/v4"
)

type App struct {
	cq  *sql.DB
	kq  *sql.DB
	env *env.Config
}

func New(env *env.Config, cq *sql.DB, kq *sql.DB) *App {
	return &App{
		cq:  cq,
		kq:  kq,
		env: env,
	}
}

func (app *App) EmptyHandlerObj(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{})
}

func (app *App) EmptyHandlerArr(c echo.Context) error {
	return c.JSON(http.StatusOK, []echo.Map{})
}
