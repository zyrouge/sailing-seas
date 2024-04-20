package routes

import (
	"net/http"
	"sailing-seas/core"

	"github.com/rs/zerolog/log"
)

func LoginRoute(app *core.App, mux *http.ServeMux) {
	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		if IsAuthenticated(r) {
			redirect(w, r)
			return
		}
		query := r.URL.Query()
		username := query.Get("username")
		password := query.Get("password")
		loginFailed := false
		if username != "" && password != "" {
			if Authenticate(app, w, username, password) {
				redirect(w, r)
				return
			}
			loginFailed = true
		}
		err := ExecuteTemplate(w, "login", map[string]any{
			"LoginFailed": loginFailed,
		})
		if err != nil {
			log.Error().Err(err).Msg("failed to render /login")
		}
	})
}
