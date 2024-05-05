package pkg

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type AssetsHandler struct {
	filename  string
	publicDir embed.FS
}

func NewAssetsHandler(publicDir embed.FS, filename string) *AssetsHandler {
	return &AssetsHandler{
		publicDir: publicDir,
		filename:  filename,
	}
}

func (h *AssetsHandler) GetPlain(c echo.Context) error {
	filename := fmt.Sprintf("%s/%s", distDir, h.filename)
	content, err := h.publicDir.ReadFile(filename)
	if err != nil {
		return c.String(http.StatusOK, os.Getenv("VERSION"))
	}
	return ResponsePlain(c, content, "0")
}

func (h *AssetsHandler) GetHTML(c echo.Context) error {
	filename := fmt.Sprintf("%s/%s", distDir, h.filename)
	content, err := h.publicDir.ReadFile(filename)
	if err != nil {
		return c.String(http.StatusOK, os.Getenv("VERSION"))
	}
	return ResponseHTML(c, content, "0")
}
