package main

import (
	"github.com/labstack/echo"
	"github.com/maxstreese/swaguiutils/pkg/swaguihandler"
)

func main() {
	swagUiHandler := swaguihandler.New(
		"https://petstore.swagger.io/v2/swagger.json", false)

	e := echo.New()
	e.GET("/*", swagUiHandler.ServeEcho)
	e.Start(":8080")
}
