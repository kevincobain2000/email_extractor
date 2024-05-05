package pkg

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type ExtractHandler struct {
	Extractor *Extract
}

func NewExtractHandler() *ExtractHandler {
	return &ExtractHandler{
		Extractor: NewExtract(),
	}
}

type ExtractorRequest struct {
	URL string `json:"url"  query:"url" validate:"required" message:"url is required"`
}

func (h *ExtractHandler) Get(c echo.Context) error {
	req := new(ExtractorRequest)
	if err := BindRequest(c, req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	if !IsURL(req.URL) {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "url is not valid")
	}
	defaults.SetDefaults(req)
	msgs, err := ValidateRequest(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	opts := []Option{
		func(o *Options) error {
			o.TimeoutMillisecond = 1000
			o.SleepMillisecond = 0
			o.URL = req.URL
			o.IgnoreQueries = true
			o.Depth = -1
			o.LimitUrls = 100
			o.LimitEmails = 1000
			o.WriteToFile = ""
			return nil
		},
	}
	enc := json.NewEncoder(c.Response())

	hc := NewHTTPChallenge(opts...)

	hc.CrawlRecursiveStream(req.URL, c, enc)

	return nil
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
