package routes

import (
	"embed"
	"net/http"
	"sailing-seas/core"
)

//go:embed static/*
var staticDir embed.FS

func StaticRoute(app *core.App, mux *http.ServeMux) {
	fileServer := http.FileServer(http.FS(staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
}
