package swaguihandler

import (
	"github.com/labstack/echo"
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
