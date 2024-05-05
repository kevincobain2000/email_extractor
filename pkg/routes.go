package pkg

import (
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	DIST_DIR     = "frontend/dist"
	FAVICON_FILE = "favicon.ico"
)

func SetupRoutes(e *echo.Echo, baseURL string, publicDir embed.FS) {
	e.GET(baseURL+"api/extract", NewExtractHandler().Get, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))
}
