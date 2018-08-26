package swaguihandler

import (
	"fmt"
	"github.com/maxstreese/swaguiutils/pkg/swaguidist"
	"net/http"
	"path/filepath"
)

func (h SwagUiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !IsValidUrlPath(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	toServeFileName := getToServeFileName(r.URL.Path)
	if toServeFileName == "" {
		toServeFileName = swaguidist.IndexFileName
	}

	setContentHeader(w, toServeFileName)

	if toServeFileName == swaguidist.IndexFileName {
		serveIndexHtml(w, h.docUrl, h.hideTopbar)
		return
	}

	fileBin, ok := swaguidist.StaticFiles[toServeFileName]
	if !ok {
		panic(fmt.Sprintf("unexpected file name: %s", toServeFileName))
	}
	w.Write(fileBin)
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
	err := swaguidist.ExecuteIndexHtml(w, docUrl, hideTopbar)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("failed to generate %s", swaguidist.IndexFileName),
			http.StatusInternalServerError)
	}
}
