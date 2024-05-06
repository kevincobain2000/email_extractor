package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const ()

type Request struct {
}

func NewRequest() *Request {
	return &Request{}
}

func SetHeaderResponseText(header http.Header) {
	header.Set("Cache-Control", "max-age=0")
	header.Set("Expires", "0")
	header.Set("Content-Type", "text/plain")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// content policy
	header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}
func SetHeadersResponsePNG(header http.Header) {
	header.Set("Cache-Control", "max-age=0")
	header.Set("Expires", "0")
	header.Set("Content-Type", "image/png")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// content policy
	header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}
func SetHeadersResponseSvg(header http.Header) {
	header.Set("Cache-Control", "max-age=0")
	header.Set("Expires", "0")
	header.Set("Content-Type", "image/svg+xml")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// // content policy
	// header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}
func SetHeadersResponseJSON(header http.Header) {
	header.Set("Cache-Control", "max-age=0")
	header.Set("Expires", "0")
	header.Set("Content-Type", "application/json")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// content policy
	header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}

func SetHeadersResponseHTML(header http.Header, cacheMS string) {
	header.Set("Cache-Control", "max-age="+cacheMS)
	header.Set("Expires", cacheMS)
	header.Set("Content-Type", "text/html; charset=utf-8")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
}

func SetHeadersResponsePlainText(header http.Header, cacheMS string) {
	header.Set("Cache-Control", "max-age="+cacheMS)
	header.Set("Expires", cacheMS)
	header.Set("Content-Type", "text/plain")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// content policy
	header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}

func ResponseMedia(c echo.Context, b []byte, media string) error {
	if media == "svg" {
		return ResponseSVG(c, b)
	}
	if media == "png" {
		return ResponsePNG(c, b)
	}
	return ResponsePNG(c, b)
}
func ResponseSVG(c echo.Context, b []byte) error {
	SetHeadersResponseSvg(c.Response().Header())
	return c.Blob(http.StatusOK, "image/svg+xml", b)
}

func ResponsePNG(c echo.Context, b []byte) error {
	SetHeadersResponsePNG(c.Response().Header())
	return c.Blob(http.StatusOK, "image/png", b)
}

func ResponseHTML(c echo.Context, b []byte, cacheMS string) error {
	SetHeadersResponseHTML(c.Response().Header(), cacheMS)
	return c.Blob(http.StatusOK, "text/html", b)
}
func ResponseIco(c echo.Context, b []byte, cacheMS string) error {
	SetHeadersResponseHTML(c.Response().Header(), cacheMS)
	return c.Blob(http.StatusOK, "image/x-icon", b)
}
func ResponsePlain(c echo.Context, b []byte, cacheMS string) error {
	SetHeadersResponsePlainText(c.Response().Header(), cacheMS)
	return c.Blob(http.StatusOK, "text/plain", b)
}
