package swaguihandler

import (
	"fmt"
	"github.com/maxstreese/swaguiutils/pkg/swaguidist"
	"strings"
)

var Paths []string

func init() {
	Paths = append(Paths, "/")
	for k, _ := range swaguidist.StaticFiles {
		Paths = append(Paths, fmt.Sprintf("/%s", k))
	}
}

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
