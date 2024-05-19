package user

import (
	"github.com/kodekoding/phastos/v2/go/api"
	"godem/src/usecase/user"
)

type Handlers interface {
	api.Controllers
}

type Handler struct {
	master Masters
	auth   Auths
}

func New(userUc user.Usecases) *Handler {
	return &Handler{master: newMaster(userUc.Master()), auth: newAuth(userUc.Auth())}
}

func (h *Handler) Register() []api.Controller {
	return []api.Controller{h.master, h.auth}
}
