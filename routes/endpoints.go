package routes

import (
	"github.com/daugvinasr/kobomatic/app"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, app *app.App) {
	e.GET("/v1/deals", app.GetDeals)
	e.GET("/v1/affiliate", app.EmptyHandlerObj)

	e.GET("/v1/user/profile", app.GetUserProfile)
	e.GET("/v1/user/wishlist", app.GetUserWishlist)
	e.GET("/v1/user/recommendations", app.GetUserRecommendations)
	e.GET("/v1/user/recommendations/feedback", app.EmptyHandlerObj)
	e.GET("/v1/user/reviews", app.GetUserReviews)
	e.GET("/v1/user/loyalty/benefits", app.GetUserLoyaltyBenefits)
	e.GET("/v1/user/browsehistory", app.EmptyHandlerArr)

	e.GET("/v1/book-images/:bookUUID/:width/:height/:quality/:isGreyscale/image.jpg", app.GetImage)
	e.GET("/v1/book-images/:bookUUID/:width/:height/:randomBoolean/image.jpg", app.GetImage)

	e.GET("/v1/initialization", app.GetInitialization)

	e.POST("/v1/analytics/gettests", app.PostAnalyticsTests)
	e.POST("/v1/analytics/event", app.EmptyHandlerObj)

	e.GET("/v1/products/dailydeal", app.EmptyHandlerObj)
	e.GET("/v1/products/featured/", app.GetFeaturedProductsTabs)
	e.GET("/v1/products/featured/:uuid", app.GetFeaturedProductsTab)
	e.GET("/v1/products/books/:uuid", app.EmptyHandlerObj)
	e.GET("/v1/products/books/series/:uuid", app.GetBookSeries)
	e.GET("/v1/products/:bookUUID/nextread", app.EmptyHandlerObj)
	e.GET("/v1/products/:bookUUID/reviews", app.GetProductReviews)

	e.POST("/v1/products/:bookUUID/reviews", app.EmptyHandlerObj)

	e.GET("/v1/products/:bookUUID/prices", app.GetProductPrice)
	e.GET("/v1/products/:bookUUID/recommendations", app.GetProductRecommendations)
	e.GET("/v1/products/:bookUUID/rating/:rating", app.EmptyHandlerObj)

	e.GET("/v1/products", app.EmptyHandlerObj)

	e.GET("/v1/categories/:categoryUUID", app.GetCategory)

	e.GET("/v1/download/:bookUUID", app.GetBookFile)

	e.GET("/v1/library/sync", app.GetSync)
	e.GET("/v1/library/:bookUUID/metadata", app.GetBookMetadata)
	e.GET("/v1/library/:bookUUID/state", app.GetBookReadingState)
	e.PUT("/v1/library/:bookUUID/state", app.PutBookReadingState)

	e.DELETE("/v1/library/:bookUUID", app.EmptyHandlerObj)
	e.DELETE("/v1/library/:bookUUID/state", app.EmptyHandlerObj)
}
