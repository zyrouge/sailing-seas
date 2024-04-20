package core

import (
	"fmt"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Route func(app *App, mux *http.ServeMux)

func StartServer(app *App, routes []Route) error {
	addr := fmt.Sprintf("%s:%d", app.Env.Host, app.Env.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Info().Msgf("listening at %s", addr)
	mux := http.NewServeMux()
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		route(app, mux)
	}
	return http.Serve(listener, mux)
}
