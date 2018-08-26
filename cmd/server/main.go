package main

import (
	"net/http"
	"github.com/maxstreese/swaggeruihandler/pkg/httphandler"
)

func main() {
	openApiDocUrl := "https://petstore.swagger.io/v2/swagger.json"
	swaggerUiHandler := httphandler.New(openApiDocUrl, false)

	http.Handle("/api/", swaggerUiHandler)

	http.ListenAndServe(":8080", nil)
}
