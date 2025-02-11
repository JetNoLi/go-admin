package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/jetnoli/go-admin/config"
	"github.com/jetnoli/go-admin/db"
	"github.com/jetnoli/go-admin/handlers"
	"github.com/jetnoli/go-admin/middleware"
	"github.com/jetnoli/go-admin/routes"
	Router "github.com/jetnoli/go-router/router"
	"github.com/jetnoli/go-router/utils"
)

func main() {

	isPortDefined, err := regexp.MatchString("^[0-9]{1,45}$", config.Port)

	if err != nil {
		fmt.Println("Port Env Variable Cannot Be Parsed")
		panic(err.Error())
	}

	utils.Assert(isPortDefined, "Port is not defined")

	db.Connect()
	defer db.Db.Close()

	router := Router.CreateRouter("*", Router.RouterOptions{
		PreHandlerMiddleware: []Router.MiddlewareHandler{middleware.DecodeToken},
	})

	router.Handle("/user/", routes.UserRouter())
	router.Handle("/auth/", routes.AuthRouter())
	router.Handle("/", routes.HTMLRouter())
	router.HandleFunc("/health/{$}", handlers.HealthCheck, &Router.RouteOptions{})

	server := http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:      router.Mux,
	}

	fmt.Println("Starting Server on http://localhost:" + config.Port)

	log.Fatal(server.ListenAndServe())
}
