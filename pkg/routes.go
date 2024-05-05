package pkg

import (
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	distDir     = "frontend/dist"
	faviconFile = "favicon.ico"
)

func SetupRoutes(e *echo.Echo, baseURL string, publicDir embed.FS) {
	e.GET(baseURL+"", NewAssetsHandler(publicDir, "index.html").GetHTML)
	e.GET(baseURL+"extract", NewAssetsHandler(publicDir, "extract/index.html").GetHTML)
	e.GET(baseURL+"api/extract", NewExtractHandler().Get, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))
}
