package swaguihandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h SwagUiHandler) ServeEcho(c echo.Context) error {
	urlPath := c.Request().URL.Path
	if !IsValidUrlPath(urlPath) {
		return echo.NewHTTPError(http.StatusNotFound, "404 Not Found")
	}

	h.ServeHTTP(c.Response().Writer, c.Request())

	return nil
}

func (h SwagUiHandler) WireUpPaths(e *echo.Echo) {
	for _, path := range Paths {
		e.GET(path, h.ServeEcho)
	}
}
