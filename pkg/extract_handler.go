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
	queue     int
}

const limitQueue = 20

func NewExtractHandler() *ExtractHandler {
	return &ExtractHandler{
		Extractor: NewExtract(),
		queue:     1,
	}
}

type ExtractorRequest struct {
	URL           string `json:"url"  query:"url" validate:"required" message:"url is required"`
	Depth         int    `json:"depth" query:"depth" default:"1" validate:"required,numeric,gte=-1,lte=5" message:"depth must be a number between -1 and 20"`
	IgnoreQueries string `json:"ignore_queries" default:"true" query:"ignore_queries" validate:"required,eq=true|eq=false" message:"ignore_queries must be true or false"`
	LimitUrls     int    `json:"limit_urls" query:"limit_urls" default:"100" validate:"required,numeric,gte=1,lte=1000" message:"limit_urls must be a number between 1 and 1000"`
	LimitEmails   int    `json:"limit_emails" query:"limit_emails" default:"1000" validate:"required,numeric,gte=1,lte=10000" message:"limit_emails must be a number between 1 and 10000"`
}

func (h *ExtractHandler) Get(c echo.Context) error {
	h.queue++
	if h.queue > limitQueue {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "queue is full, please try again later")
	}

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

	opts := []CrawlOption{
		func(o *CrawlOptions) error {
			o.TimeoutMillisecond = 1000
			o.SleepMillisecond = 10
			o.URL = req.URL
			o.IgnoreQueries = req.IgnoreQueries == "true"
			o.Depth = req.Depth
			o.LimitUrls = req.LimitUrls
			o.LimitEmails = req.LimitEmails
			o.WriteToFile = ""
			return nil
		},
	}
	enc := json.NewEncoder(c.Response())

	hc := NewHTTPChallenge(opts...)

	hc.CrawlRecursiveStream(req.URL, c, enc)

	h.queue--

	return nil
}

func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
