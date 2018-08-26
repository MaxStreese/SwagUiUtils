package httphandler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"github.com/maxstreese/swaggeruihandler/pkg/dist"
)

type SwaggerUiHandler struct {
	docUrl     string
	hideTopbar bool
}

func New(docUrl string, hideTopbar bool) SwaggerUiHandler {
	return SwaggerUiHandler{docUrl, hideTopbar}
}

func (h SwaggerUiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !IsValidUrlPath(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	toServeFileName := getToServeFileName(r.URL.Path)
	if toServeFileName == "" {
		toServeFileName = "index.html"
	}

	setContentHeader(w, toServeFileName)

	if toServeFileName == "index.html" {
		serveIndexHtml(w, h.docUrl, h.hideTopbar)
		return
	}

	fileBin, ok := dist.StaticFiles[toServeFileName]
	if !ok {
		panic(fmt.Sprintf("unexpected file name: %s", toServeFileName))
	}
	w.Write(fileBin)
}

func IsValidUrlPath(urlPath string) bool {
	toServeFileName := getToServeFileName(urlPath)

	if toServeFileName == "" {
		return true
	}

	_, ok := dist.StaticFiles[toServeFileName]
	return ok
}

func getToServeFileName(urlPath string) string {
	urlPathParts := strings.Split(urlPath, "/")
	return urlPathParts[len(urlPathParts)-1]
}

func setContentHeader(w http.ResponseWriter, fileName string) {
	headerKey := "Content-Type"
	fileExtension := filepath.Ext(fileName)
	switch fileExtension {
	case ".html":
		w.Header().Set(headerKey, "text/html")
	case ".js":
		w.Header().Set(headerKey, "application/javascript")
	case ".css":
		w.Header().Set(headerKey, "text/css")
	case ".png":
		w.Header().Set(headerKey, "application/octet-stream")
	default:
		panic(fmt.Sprintf("unexpected file extension: %s", fileExtension))
	}
}

func serveIndexHtml(w http.ResponseWriter, docUrl string, hideTopbar bool) {
	err := dist.ExecuteIndexHtml(w, docUrl, hideTopbar)
	if err != nil {
		http.Error(w, "failed to generate index.html",
			http.StatusInternalServerError)
	}
}
