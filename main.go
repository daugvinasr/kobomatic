package main

import (
	"database/sql"
	"embed"

	"github.com/daugvinasr/kobomatic/app"
	"github.com/daugvinasr/kobomatic/env"
	"github.com/daugvinasr/kobomatic/middleware"
	"github.com/daugvinasr/kobomatic/routes"
	"github.com/labstack/echo/v4"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	env, err := env.Load()
	if err != nil {
		logrus.Fatal("failed to load configuration:", err)
	}

	calibreDB, err := sql.Open("sqlite", env.LibraryFolder+"/metadata.db?mode=ro")
	if err != nil {
		logrus.Fatal("failed to open calibre database", err)
	}
	defer calibreDB.Close()

	kobomaticDB, err := sql.Open("sqlite", env.KobomaticFolder+"/kobomatic.db")
	if err != nil {
		logrus.Fatal("failed to open kobomatic database", err)
	}
	defer kobomaticDB.Close()

	goose.SetLogger(logrus.StandardLogger())
	goose.SetBaseFS(embedMigrations)

	err = goose.SetDialect("sqlite")
	if err != nil {
		logrus.Fatal("failed to set goose dialect", err)
	}

	err = goose.Up(kobomaticDB, "migrations")
	if err != nil {
		logrus.Fatal("failed to migrate kobomatic database", err)
	}

	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.LogrusLogger)

	a := app.New(env, calibreDB, kobomaticDB)

	routes.SetupRoutes(e, a)

	logrus.Error(e.Start(":8084"))
}
