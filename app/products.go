package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *App) GetProductPrice(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items": []string{},
	})
}

func (app *App) GetProductRecommendations(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"ItemCount":        0,
		"TotalPageCount":   0,
		"TotalItemCount":   0,
		"CurrentPageIndex": 0,
		"ItemsPerPage":     100,
	})
}

func (app *App) GetProductReviews(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"TotalPageCount":   0,
		"CurrentPageIndex": 0,
	})
}

func (app *App) GetFeaturedProductsTabs(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"ItemCount":        0,
		"TotalPageCount":   0,
		"TotalItemCount":   0,
		"CurrentPageIndex": 0,
		"ItemsPerPage":     100,
	})
}

func (app *App) GetFeaturedProductsTab(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Items":            []string{},
		"ItemCount":        0,
		"TotalPageCount":   0,
		"TotalItemCount":   0,
		"CurrentPageIndex": 0,
		"ItemsPerPage":     100,
		"Filters":          echo.Map{"PreOrders": []string{"True", "False"}},
	})
}

func (app *App) GetBookSeries(c echo.Context) error {
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
