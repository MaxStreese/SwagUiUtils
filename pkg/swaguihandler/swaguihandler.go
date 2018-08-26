package swaguihandler

import (
	"github.com/maxstreese/swaguiutils/pkg/swaguidist"
	"strings"
)

type SwagUiHandler struct {
	docUrl     string
	hideTopbar bool
}

func New(docUrl string, hideTopbar bool) SwagUiHandler {
	return SwagUiHandler{docUrl, hideTopbar}
}

func IsValidUrlPath(urlPath string) bool {
	toServeFileName := getToServeFileName(urlPath)

	if toServeFileName == "" {
		return true
	}

	_, ok := swaguidist.StaticFiles[toServeFileName]
	return ok
}

func getToServeFileName(urlPath string) string {
	urlPathParts := strings.Split(urlPath, "/")
	return urlPathParts[len(urlPathParts)-1]
}
