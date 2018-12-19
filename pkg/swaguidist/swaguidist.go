package swaguidist

import (
	"html/template"
	"io"
)

var (
	IndexFileName = "index.html"

	StaticFiles = map[string][]byte{
		"favicon-16x16.png":               favicon16x16Png,
		"favicon-32x32.png":               favicon32x32Png,
		"oauth2-redirect.html":            oAuth2RedirectHtml,
		"swagger-ui-bundle.js":            swaggerUiBundleJs,
		"swagger-ui.css":                  swaggerUiCss,
		"swagger-ui-standalone-preset.js": swaggerUiStandalonePresetJs,
	}

	indexHtmlTemplate = template.Must(
		template.New(IndexFileName).Parse(indexHtml))

	pluginDownloadUrl = "SwaggerUIBundle.plugins.DownloadUrl"
	pluginHideTopbar  = "HideTopbarPlugin"
)

func ExecuteIndexHtml(w io.Writer, docUrl string, hideTopbar bool) error {
	plugins := pluginDownloadUrl
	if hideTopbar {
		plugins = plugins + ", " + pluginHideTopbar
	}

	config := struct {
		DocUrl  string
		Plugins template.JS
	}{
		DocUrl:  docUrl,
		Plugins: template.JS(plugins),
	}

	return indexHtmlTemplate.Execute(w, config)
}
