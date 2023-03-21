package api

import (
	"github.com/kodekoding/phastos/go/apps"
	"github.com/kodekoding/phastos/go/notifications"
	"github.com/kodekoding/phastos/go/response"
	"godem/infrastructure/middleware"
	"godem/infrastructure/middleware/auth"
	"godem/infrastructure/service/_internal/http/api/v1/user"
	"net/http"
)

type (
	Handlers interface {
		Modules() *Modules
		Middleware() *middleware.Definitions
		Wrapper() *WrapperHandler
		Ping(_ http.ResponseWriter, _ *http.Request) *response.JSON
	}
	Handler struct {
		modules     *Modules
		HttpHandler http.Handler
		wrapper     *WrapperHandler
		middleware  *middleware.Definitions
	}
	WrapperHandler struct {
		Notif notifications.Platforms
		Apps  apps.Slacks
	}
	Modules struct {
		User user.Handlers
		Auth auth.Middlewares
	}
)

func NewHandler(modules *Modules, wrapper *WrapperHandler, middle *middleware.Definitions) *Handler {

	return &Handler{modules: modules, wrapper: wrapper, middleware: middle}
}

func (h *Handler) Modules() *Modules {
	return h.modules
}

func (h *Handler) Wrapper() *WrapperHandler {
	return h.wrapper
}

func (h *Handler) Middleware() *middleware.Definitions {
	return h.middleware
}

func (h *Handler) Ping(_ http.ResponseWriter, _ *http.Request) *response.JSON {
	return response.NewJSON().Success("PONG")
}
