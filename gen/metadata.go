package gen

import (
	"fmt"
	"strconv"

	"github.com/daugvinasr/kobomatic/database/calibre"
	"github.com/daugvinasr/kobomatic/helpers"
	"github.com/labstack/echo/v4"
)

func Metadata(serverAddress string, bm calibre.FindBookMetadataRow) echo.Map {
	crossRevisionID := helpers.GetDeterministicUUID(bm.UUID + "crossRevisionID")
	revisionID := helpers.GetDeterministicUUID(bm.UUID + "revisionID")
	relatedGroupID := helpers.GetDeterministicUUID(bm.UUID + "relatedGroupID")

	metadata := echo.Map{
		"CrossRevisionId": crossRevisionID,
		"RevisionId":      revisionID,
		"Publisher": echo.Map{
			"Name":    bm.Publisher,
			"Imprint": "",
		},
		"PublicationDate":         "2024-01-01T00:00:00.0000000",
		"Language":                "en",
		"Isbn":                    "0000000000001",
		"Genre":                   "00000000-0000-0000-0000-000000000001",
		"Slug":                    bm.UUID,
		"CoverImageId":            bm.UUID,
		"ApplicableSubscriptions": []string{},
		"IsSocialEnabled":         true,
		"WorkId":                  crossRevisionID,
		"ExternalIds":             []string{},
		"IsPreOrder":              false,
		"ContributorRoles": []echo.Map{
			{"Name": bm.Author},
		},
		"IsInternetArchive":          false,
		"IsAnnotationExportDisabled": false,
		"IsAISummaryDisabled":        false,
		"EntitlementId":              bm.UUID,
		"Title":                      bm.Title,
		"Description":                bm.Description,
		"Categories":                 []string{},
		"DownloadUrls": []echo.Map{
			{
				"DrmType":  "None",
				"Format":   "KEPUB",
				"Url":      fmt.Sprintf("%s/v1/download/%s", serverAddress, bm.UUID),
				"Platform": "Generic",
				"Size":     bm.Size,
			},
		},
		"Contributors": []string{bm.Author},
		"CurrentDisplayPrice": echo.Map{
			"TotalAmount":  0,
			"CurrencyCode": "EUR",
		},
		"CurrentLoveDisplayPrice": echo.Map{
			"TotalAmount": 0,
		},
		"IsEligibleForKoboLove":  false,
		"PhoneticPronunciations": echo.Map{},
		"RelatedGroupId":         relatedGroupID,
	}

	if bm.Series != nil && bm.SeriesIndex != nil {
		metadata["Series"] = echo.Map{
			"Id":          helpers.GetDeterministicUUID(*bm.Series).String(),
			"Name":        *bm.Series,
			"Number":      strconv.FormatFloat(*bm.SeriesIndex, 'f', -1, 64),
			"NumberFloat": *bm.SeriesIndex,
		}
	}

	return metadata
}
