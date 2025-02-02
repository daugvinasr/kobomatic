package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) GetUserProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"IsOneStore":              true,
		"IsChildAccount":          false,
		"CountryCode":             "LT",
		"Geo":                     "LT",
		"StoreFront":              "LT",
		"PlatformId":              "00000000-0000-0000-0000-000000000000",
		"PartnerId":               "00000000-0000-0000-0000-000000000001",
		"AffiliateName":           "Kobo",
		"IsoCultureCode":          "en-US",
		"IsLibraryMigrated":       false,
		"VipMembershipPurchased":  false,
		"HasPurchased":            false,
		"HasPurchasedBook":        false,
		"HasPurchasedAudiobook":   false,
		"SafeSearch":              false,
		"AudiobooksEnabled":       true,
		"IsOrangeAffiliated":      false,
		"IsEligibleForOrangeDeal": false,
		"PrivacyPermissions":      []string{},
	})
}

func (app *App) GetUserWishlist(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"TotalCountByProductType": echo.Map{},
		"Items":                   []string{},
		"ItemCount":               0,
		"TotalPageCount":          0,
		"TotalItemCount":          0,
		"CurrentPageIndex":        0,
		"ItemsPerPage":            100,
	})
}

func (app *App) GetUserRecommendations(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"ItemCount":        0,
		"TotalPageCount":   0,
		"TotalItemCount":   0,
		"CurrentPageIndex": 0,
		"ItemsPerPage":     100,
		"Filters":          echo.Map{},
	})
}

func (app *App) GetUserReviews(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"TotalPageCount":   0,
		"CurrentPageIndex": 0,
	})
}

func (app *App) GetUserLoyaltyBenefits(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Benefits": echo.Map{},
	})
}
