package routes

import (
	"io"
	"net/http"
	"sailing-seas/core"
)

func PingRoute(app *core.App, mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong!\n")
	})
}
