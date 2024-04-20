package routes

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/maps"
)

var cachedTemplates *template.Template

func ExecuteTemplate(wr io.Writer, name string, data map[string]any) error {
	if cachedTemplates == nil {
		files, err := getTemplateFiles()
		if err != nil {
			return err
		}
		parsed, err := template.New("_").Funcs(getTemplateFuncMap()).ParseFiles(files...)
		if err != nil {
			return err
		}
		cachedTemplates = parsed
	}
	data["SiteName"] = "Sailing Seas"
	return cachedTemplates.ExecuteTemplate(wr, name, data)
}

func getTemplateFiles() ([]string, error) {
	dirs := []string{"routes/templates/*.gohtml", "routes/templates/partials/*.gohtml"}
	files := []string{}
	for _, x := range dirs {
		matched, err := filepath.Glob(x)
		if err != nil {
			return nil, err
		}
		files = append(files, matched...)
	}
	return files, nil
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
