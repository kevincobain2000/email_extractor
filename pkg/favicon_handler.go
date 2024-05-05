package pkg

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FaviconHandler struct {
	favicon  *embed.FS
	filename string
}

func NewFaviconHandler(favicon *embed.FS) *FaviconHandler {
	return &FaviconHandler{
		favicon:  favicon,
		filename: FAVICON_FILE,
	}
}

func (h *FaviconHandler) Get(c echo.Context) error {
	SetHeadersResponsePNG(c.Response().Header())
	content, err := h.favicon.ReadFile(h.filename)
	if err != nil {
		return err
	}
	SetHeadersResponsePNG(c.Response().Header())
	return c.Blob(http.StatusOK, "image/x-icon", content)
}
