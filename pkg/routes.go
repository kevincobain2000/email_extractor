package pkg

import (
	"github.com/labstack/echo/v4"
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
	// no longer providing it as a service
	// e.GET(opt.BaseURL+"extract", NewAssetsHandler(opt.PublicDir, "extract/index.html").GetHTML)

	e.GET(opt.BaseURL+"extract", func(c echo.Context) error {
		return c.Redirect(301, "https://github.com/kevincobain2000/email_extractor")
	})
	e.GET(opt.BaseURL+"ads.txt", NewAssetsHandler(opt.PublicDir, "ads.txt").GetPlain)
	e.GET(opt.BaseURL+"robots.txt", NewAssetsHandler(opt.PublicDir, "robots.txt").GetPlain)
	e.GET(opt.BaseURL+"favicon.ico", NewAssetsHandler(opt.PublicDir, "favicon.ico").GetIco)
	// no longer providing it as a service
	// e.GET(opt.BaseURL+"api/extract",
	// 	NewExtractHandler().Get,
	// 	middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)),
	// )
	e.GET(opt.BaseURL+"api/extract", func(c echo.Context) error {
		return c.Redirect(301, "https://github.com/kevincobain2000/email_extractor")
	})
}
