package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kodekoding/phastos/go/router"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type (
	Routes interface {
		Register()
	}
	Route struct {
		handler      Handlers
		handle       router.RouteInterface
		listOfRoutes map[string][]string
	}
)

func NewRoutes(handlers Handlers) *Route {
	return &Route{handle: router.NewChiRouter(), handler: handlers, listOfRoutes: make(map[string][]string)}
}

func (r *Route) GetHandler() *chi.Mux {
	return r.handle.GetHandler()
}

func (r *Route) Register() {
	route := r.handle
	//tracer, closer := jaeger.NewTracer(
	//	"serviceName",
	//	jaeger.NewConstSampler(true),
	//	jaeger.NewInMemoryReporter(),
	//)
	//defer closer.Close()
	route.InitRoute(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"PATCH", "POST", "DELETE", "GET", "PUT"},
		AllowedHeaders: []string{"Origin", "token", "content-type", "Content-Type", "Authorization"},
		MaxAge:         60 * 60, //1 hour
	})

	//buat url image setelah di upload
	//route.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	//handler := r.handler.Modules()
	//swagger url
	//route.Get("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route.Get("/ping", r.handler.Ping)
	const prefix string = "/api"

	go func() {
		r.trackAllRoutes()
	}()
}
func (r *Route) appendRoutes(method, route string) {
	r.listOfRoutes[route] = append(r.listOfRoutes[route], method)
}

func (r *Route) trackAllRoutes() {
	endpointCount := 0
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		r.appendRoutes(method, route)
		endpointCount++
		return nil
	}

	if err := chi.Walk(r.handle.GetHandler(), walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
	log.Printf("Served %d endpoints", endpointCount)

}
