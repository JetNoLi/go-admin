package users

import (
	Router "github.com/jetnoli/go-router/router"
)

func userRouter() *Router.Router {
	router := Router.CreateRouter("/users", Router.RouterOptions{
		ExactPathsOnly: false,
	})

	router.Post("/signup", SignUp, &Router.RouteOptions{})
	router.Get("/", GetAllUsers, &Router.RouteOptions{})
	router.Get("/{id}", GetUserById, &Router.RouteOptions{})
	router.Get("/username/{username}", GetUserByUsername, &Router.RouteOptions{})
	router.Put("/{id}", UpdateUserById, &Router.RouteOptions{})
	router.Delete("/{id}", DeleteUserById, &Router.RouteOptions{})
	router.Delete("/", DeleteAllUsers, &Router.RouteOptions{})

	return router
}

var UserRouter = userRouter()
