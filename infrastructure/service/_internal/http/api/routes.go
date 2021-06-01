package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"godem/infrastructure/service/_internal/http/api/user"
	"godem/lib/util/response"
)

type Routes struct {
	router  *chi.Mux
	modules *ModuleHandler
}

type ModuleHandler struct {
	User *user.Handler
}

func NewHandler(modules *ModuleHandler) *Routes {
	return &Routes{
		router:  chi.NewRouter(),
		modules: modules,
	}
}

func (h *Routes) RegisterAndStartServer() error {

	//register your routes here
	h.router.Get("/ping", h.Ping)

	h.router.Route("/user", func(user chi.Router) {
		user.Get("/", h.modules.User.Master().GetList)
		user.Get("/{id}", h.modules.User.Master().GetDetailByID)
		user.Post("/", h.modules.User.Master().CreateNew)
		user.Patch("/{id}", h.modules.User.Master().UpdateData)
		user.Delete("/{id}", h.modules.User.Master().DeleteData)

	})

	log.Println("http listening on port :10000")
	return http.ListenAndServe(":10000", h.router)
}

func (h *Routes) Ping(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"code":    200,
		"message": "OK",
	}

	resp := response.NewJSON()
	resp.Success(data).Send(w)
}
