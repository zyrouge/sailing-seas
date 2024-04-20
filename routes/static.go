package routes

import (
	"net/http"
	"path/filepath"
	"sailing-seas/core"
)

func StaticRoute(app *core.App, mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir(filepath.Join("routes/static")))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
}
