package routes

import (
	"net/http"
	"sailing-seas/core"
	"sailing-seas/helpers"

	"github.com/rs/zerolog/log"
)

func NyaaRoute(app *core.App, mux *http.ServeMux) {
	mux.HandleFunc("GET /nyaa", NeedsAuthentication(nyaaRouteHandler()))
}

func nyaaRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		terms := r.URL.Query().Get("q")
		items, err := helpers.NyaaSearch(terms)
		if err != nil {
			log.Error().Err(err).Msgf("nyaa search failed for query '%s'", terms)
			items = []helpers.NyaaSearchItem{}
		}
		err = ExecuteTemplate(w, "nyaa", map[string]any{
			"HasSearchTerms":   terms != "",
			"SearchTerms":      terms,
			"HasSearchResults": len(items) != 0,
			"SearchResults":    items,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to render /nyaa")
		}
	}
}
