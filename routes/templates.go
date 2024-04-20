package routes

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/maps"
)

var cachedTemplates *template.Template

//go:embed templates/*.gotmpl templates/partials/*.gotmpl
var templatesDir embed.FS

func ExecuteTemplate(wr io.Writer, name string, data map[string]any) error {
	if cachedTemplates == nil {
		temp := template.New("_").Funcs(getTemplateFuncMap())
		parsed, err := temp.ParseFS(templatesDir, "templates/*.gotmpl", "templates/partials/*.gotmpl")
		if err != nil {
			return err
		}
		cachedTemplates = parsed
	}
	data["SiteName"] = "Sailing Seas"
	return cachedTemplates.ExecuteTemplate(wr, name, data)
}

const templateFilterFail = "ZsstmplZ"

func getTemplateFuncMap() map[string]any {
	out := sprig.FuncMap()
	maps.Copy(out, template.FuncMap{
		"magnetUrl": func(value string) template.URL {
			if !strings.HasPrefix(value, "magnet:") {
				return "#" + templateFilterFail
			}
			return template.URL(value)
		},
	})
	return out
}
