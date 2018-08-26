package main

import (
	"net/http"
	"github.com/maxstreese/swaguiutils/pkg/swaguihandler"
)

func main() {
	openApiDocUrl := "https://petstore.swagger.io/v2/swagger.json"
	swagUiHandler := swaguihandler.New(openApiDocUrl, false)

	http.Handle("/api/", swagUiHandler)

	http.ListenAndServe(":8080", nil)
}
