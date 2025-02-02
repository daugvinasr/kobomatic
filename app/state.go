package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/daugvinasr/kobomatic/constants"
	"github.com/daugvinasr/kobomatic/database/kobomatic"
	"github.com/daugvinasr/kobomatic/gen"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (app *App) GetBookReadingState(c echo.Context) error {
	entitlementID := c.Param("bookUUID")

	rs, err := kobomatic.GetReadingState(app.kq, entitlementID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusOK, []gen.ReadingStateComplete{})
		}
		return err
	}

	crs, err := gen.CompleteReadingState(rs)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, []gen.ReadingStateComplete{crs})
}

func shouldUpdateReadingState(db *sql.DB, readingState gen.ReadingStateComplete) (bool, error) {
	storedLastModified, err := kobomatic.GetReadingStateLastModified(db, readingState.EntitlementID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}

	parsedStored, err := time.Parse(constants.ISO8601, storedLastModified)
	if err != nil {
		return false, err
	}

	// kobo sends in time in something that looks like RFC3339, but is not?
	parsedNew, err := time.Parse(constants.KoboTime, readingState.LastModified)
	if err != nil {
		return false, err
	}

	return parsedNew.After(parsedStored), nil
}

func (app *App) PutBookReadingState(c echo.Context) error {
	var requestBody struct {
		ReadingStates []gen.ReadingStateComplete `json:"ReadingStates"`
	}

	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		return err
	}

	defer c.Request().Body.Close()

	if len(requestBody.ReadingStates) == 0 {
		return errors.New("no ReadingState was provided")
	}

	readingState := requestBody.ReadingStates[0]

	shouldUpdate, err := shouldUpdateReadingState(app.kq, readingState)
	if err != nil {
		return err
	}

	if shouldUpdate {
		// kobo sends in time in something that looks like RFC3339, but is not?
		parsedLastModified, err := time.Parse(constants.KoboTime, readingState.LastModified)
		if err != nil {
			return err
		}

		lastModified := parsedLastModified.Format(constants.ISO8601)

		err = kobomatic.InsertReadingState(app.kq, kobomatic.ReadingState{
			EntitlementID:                readingState.EntitlementID,
			LastModified:                 lastModified,
			Status:                       readingState.StatusInfo.Status,
			SpentReadingMinutes:          readingState.Statistics.SpentReadingMinutes,
			RemainingTimeMinutes:         readingState.Statistics.RemainingTimeMinutes,
			ProgressPercent:              readingState.CurrentBookmark.ProgressPercent,
			ContentSourceProgressPercent: readingState.CurrentBookmark.ContentSourceProgressPercent,
			Value:                        readingState.CurrentBookmark.Location.Value,
			Type:                         readingState.CurrentBookmark.Location.Type,
			Source:                       readingState.CurrentBookmark.Location.Source,
		})
		if err != nil {
			return err
		}
	} else {
		logrus.Infof("ReadingState for entitlement %s was ignored", readingState.EntitlementID)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"RequestResult": "Success",
		"UpdateResults": []echo.Map{
			{
				"EntitlementId":         readingState.EntitlementID,
				"StatusInfoResult":      echo.Map{"Result": "Success"},
				"StatisticsResult":      echo.Map{"Result": "Success"},
				"CurrentBookmarkResult": echo.Map{"Result": "Success"},
			},
		},
	})
}
