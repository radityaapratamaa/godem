package main

import (
	"github.com/kodekoding/phastos/go/server"
	"github.com/pkg/errors"
	"github.com/unrolled/secure"
	"godem/infrastructure/service/_internal/router/component"
	"log"
	"net/http"

	"godem/infrastructure/config"
	"godem/infrastructure/loader"
	"godem/infrastructure/middleware"
	"godem/infrastructure/service/_internal/http/api"
)

func main() {
	// read config file
	cfg, err := config.New()
	if err != nil {
		log.Fatalln("cannot read config: ", errors.Cause(err).Error())
	}
	log.Printf("%#v", cfg)
	modules := loader.InitModules(cfg)
	wrappers := loader.InitWrapper(cfg)
	middlewares := middleware.New(modules.Auth)

	handler := api.NewHandler(modules, wrappers, middlewares)

	routeHandler := loadHandler(handler)
	serverConfig := cfg.Server
	serverConfig.Handler = routeHandler
	if err = server.ServeHTTP(&serverConfig); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}

func loadHandler(this *api.Handler) http.Handler {
	routes := api.NewRoutes(this)
	// register routes
	routes.Register()

	this.HttpHandler = component.InitHandler(routes.GetHandler())

	secureMiddleware := secure.New(secure.Options{
		BrowserXssFilter:   true,
		ContentTypeNosniff: true,
	})

	this.HttpHandler = secureMiddleware.Handler(this.HttpHandler)

	this.HttpHandler = component.WrapNotif(this.HttpHandler, this.Wrapper().Notif)
	return this.HttpHandler
}
