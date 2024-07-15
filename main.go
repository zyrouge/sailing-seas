package main

import (
	_ "github.com/joho/godotenv/autoload"

	"sailing-seas/core"
	"sailing-seas/routes"

	"github.com/rs/zerolog/log"
)

func main() {
	app, err := core.CreateApp()
	if err != nil {
		log.Panic().Err(err).Msg("app creation failed")
	}
	routes := []core.Route{
		routes.StaticRoute,
		routes.PingRoute,
		routes.HomeRoute,
		routes.LoginRoute,
		routes.NyaaRoute,
	}
	if err := core.StartServer(app, routes); err != nil {
		log.Panic().Err(err).Send()
	}
}
