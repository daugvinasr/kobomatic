package app

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/daugvinasr/kobomatic/constants"
	"github.com/daugvinasr/kobomatic/database/calibre"
	"github.com/daugvinasr/kobomatic/database/kobomatic"
	"github.com/daugvinasr/kobomatic/gen"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type SyncToken struct {
	Version  string `json:"version"`
	LastSync string `json:"lastSync"`
}

func generateKobomaticSyncToken(old bool) SyncToken {
	var lastSync string

	if old {
		lastSync = constants.OldTimestamp
	} else {
		lastSync = time.Now().UTC().Format(constants.ISO8601)
	}

	return SyncToken{
		Version:  "kobomatic-1.0",
		LastSync: lastSync,
	}
}

func getKobomaticSyncToken(c echo.Context) SyncToken {
	b64SyncHeader := c.Request().Header.Get("x-kobo-synctoken")
	if b64SyncHeader == "" {
		return generateKobomaticSyncToken(true)
	}

	syncTokenBytes, err := base64.StdEncoding.DecodeString(b64SyncHeader)
	if err != nil {
		return generateKobomaticSyncToken(true)
	}

	var syncToken SyncToken

	err = json.Unmarshal(syncTokenBytes, &syncToken)
	if err != nil {
		return generateKobomaticSyncToken(true)
	}

	return syncToken
}

// nolint: cyclop
func (app *App) GetSync(c echo.Context) error {
	requestSyncToken := getKobomaticSyncToken(c)

	booksMetadata, err := calibre.FindBooksMetadata(app.cq, requestSyncToken.LastSync)
	if err != nil {
		return err
	}

	e := []echo.Map{}

	syncedEntitlementMap := make(map[string]bool)

	for _, bm := range booksMetadata {
		entitlement := echo.Map{
			"BookEntitlement": gen.Entitlement(bm.UUID),
			"BookMetadata":    gen.Metadata(app.env.ServerAddress, bm),
		}

		rs, err := kobomatic.GetReadingState(app.kq, bm.UUID)
		if err != nil {
			entitlement["ReadingState"] = gen.MinimalReadingState(bm.UUID)
		} else {
			crs, err := gen.CompleteReadingState(rs)
			if err != nil {
				return err
			}

			entitlement["ReadingState"] = crs
		}

		e = append(e, echo.Map{"NewEntitlement": entitlement})
		syncedEntitlementMap[bm.UUID] = true
	}

	logrus.Infof("returned %d NewEntitlement for /v1/library/sync", len(syncedEntitlementMap))

	// we don't send reading state from other devices on initial sync
	if requestSyncToken.LastSync != constants.OldTimestamp {
		newReadingStates, err := kobomatic.GetReadingStatesByLastModified(app.kq, requestSyncToken.LastSync)
		if err != nil {
			return err
		}

		// we need to filter out reading states that would be synced with NewEntitlement block
		filteredReadingStates := []kobomatic.ReadingState{}
		for _, rs := range newReadingStates {
			if !syncedEntitlementMap[rs.EntitlementID] {
				filteredReadingStates = append(filteredReadingStates, rs)
			}
		}

		for _, rs := range filteredReadingStates {
			crs, err := gen.CompleteReadingState(rs)
			if err != nil {
				return err
			}

			e = append(e, echo.Map{"ChangedReadingState": crs})
		}

		logrus.Infof("returned %d ChangedReadingState for /v1/library/sync", len(filteredReadingStates))
	}

	responseSyncToken, err := json.Marshal(generateKobomaticSyncToken(false))
	if err != nil {
		return err
	}

	c.Response().Header().Set("x-kobo-synctoken", base64.StdEncoding.EncodeToString(responseSyncToken))

	return c.JSON(http.StatusOK, e)
}
