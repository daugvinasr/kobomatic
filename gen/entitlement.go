package gen

import (
	"github.com/daugvinasr/kobomatic/constants"
	"github.com/daugvinasr/kobomatic/helpers"
	"github.com/labstack/echo/v4"
)

func Entitlement(entitlementID string) echo.Map {
	crossRevisionID := helpers.GetDeterministicUUID(entitlementID + "crossRevisionID")
	revisionID := helpers.GetDeterministicUUID(entitlementID + "revisionID")

	return echo.Map{
		"Accessibility": "Full",
		"ActivePeriod": echo.Map{
			"From": constants.KoboOldTimestamp,
		},
		"Created":             constants.KoboOldTimestamp,
		"CrossRevisionId":     crossRevisionID,
		"Id":                  entitlementID,
		"IsRemoved":           false,
		"IsHiddenFromArchive": false,
		"IsLocked":            false,
		"LastModified":        constants.KoboOldTimestamp,
		"OriginCategory":      "Imported",
		"RevisionId":          revisionID,
		"Status":              "Active",
	}
}
