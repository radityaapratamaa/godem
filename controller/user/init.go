package user

import (
	"godem/usecase/user"
)

type Handlers interface {
	Master() Masters
	Auth() Auths
}

type Handler struct {
	master Masters
	auth   Auths
}

func New(userUc user.Usecases) *Handler {
	return &Handler{master: newMaster(userUc.Master()), auth: newAuth(userUc.Auth())}
}

func (h *Handler) Master() Masters {
	return h.master
}

func (h *Handler) Auth() Auths {
	return h.auth
}
