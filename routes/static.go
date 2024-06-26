package routes

import (
	"embed"
	"io/fs"
	"net/http"
	"sailing-seas/core"

	"github.com/rs/zerolog/log"
)

//go:embed static/*
var staticDir embed.FS

func StaticRoute(app *core.App, mux *http.ServeMux) {
	staticSubDir, err := fs.Sub(staticDir, "static")
	if err != nil {
		log.Panic().Err(err).Send()
	}
	fileServer := http.FileServerFS(staticSubDir)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}
