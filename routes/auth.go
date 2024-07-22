package routes

import (
	"net/http"

	"github.com/jetnoli/go-admin/handlers"
	Router "github.com/jetnoli/go-router/router"
)

func AuthRouter() *http.ServeMux {
	router := Router.CreateRouter("/auth", Router.RouterOptions{})

	router.Post("/login", handlers.SignInHtmx, &Router.RouteOptions{})
	router.Post("/signup", handlers.SignUpHtmx, &Router.RouteOptions{})

	return router.Mux
}
