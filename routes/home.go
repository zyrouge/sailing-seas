package routes

import (
	"net/http"
	"sailing-seas/core"

	"github.com/rs/zerolog/log"
)

func HomeRoute(app *core.App, mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", NeedsAuthentication(homeRouteHandler(app)))
}

func homeRouteHandler(app *core.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ExecuteTemplate(w, "home", map[string]any{})
		if err != nil {
			log.Error().Err(err).Msg("failed to render /")
		}
	}
}
