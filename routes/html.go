package routes

import (
	"net/http"

	"github.com/jetnoli/go-admin/middleware"

	"github.com/jetnoli/go-admin/view/pages/admin"
	"github.com/jetnoli/go-admin/view/pages/home"
	"github.com/jetnoli/go-admin/view/pages/login"
	"github.com/jetnoli/go-admin/view/pages/signup"

	Router "github.com/jetnoli/go-router/router"
)

func HTMLRouter() *http.ServeMux {
	router := Router.CreateRouter("/", Router.RouterOptions{
		ExactPathsOnly: true,
	})

	// Serve Styles for Pages
	router.ServeDir("/", "view/pages/", &Router.ServeDirOptions{
		IncludedExtensions:         []string{".css"},
		Recursive:                  true,
		RoutePathContainsExtension: true,
	})

	// Serve Styles for Components
	router.ServeDir("/", "view/components/", &Router.ServeDirOptions{
		IncludedExtensions:         []string{".css"},
		Recursive:                  true,
		RoutePathContainsExtension: true,
	})

	router.ServeDir("/assets/", "assets/", &Router.ServeDirOptions{
		Recursive:                  true,
		RoutePathContainsExtension: true,
	})

	router.ServeTempl(map[string]*Router.TemplPage{
		"/": {
			PageComponent: home.Page(),
			Options: &Router.RouteOptions{
				PreHandlerMiddleware: []Router.MiddlewareHandler{middleware.CheckAuthorization},
			},
		},
		"/login": {
			PageComponent: login.Page(),
		},
		"/signup": {
			PageComponent: signup.Page(),
		},
		"/admin/": {
			PageComponent: admin.Page(admin.PageUrlParams{}),
		},
		// "/admin/users/": {
		// 	PageComponent: admin.Page(admin.PageUrlParams{Module: module.UserItemList()}),
		// },
		// "/admin/{module}/{id}": {

		// },
	})

	return router.Mux
}
