package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, options []EchoOption) {
	opt := &EchoOptions{}
	for _, o := range options {
		err := o(opt)
		if err != nil {
			panic(err)
		}
	}
	e.GET(opt.BaseURL+"", NewAssetsHandler(opt.PublicDir, "index.html").GetHTML)
	e.GET(opt.BaseURL+"extract", NewAssetsHandler(opt.PublicDir, "extract/index.html").GetHTML)
	e.GET(opt.BaseURL+"robots.txt", NewAssetsHandler(opt.PublicDir, "robots.txt").GetPlain)
	e.GET(opt.BaseURL+"favicon.ico", NewAssetsHandler(opt.PublicDir, "favicon.ico").GetIco)
	e.GET(opt.BaseURL+"api/extract",
		NewExtractHandler().Get,
		middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)),
	)
}
