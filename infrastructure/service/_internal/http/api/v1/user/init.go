package user

import "godem/usecase/user"

type Handlers interface {
	Master() Masters
	Login() Logins
}

type Handler struct {
	master Masters
	login  Logins
}

func New(userUc user.Usecases) *Handler {
	return &Handler{master: newMaster(userUc.Master()), login: newLogin(userUc.Login())}
}

func (h *Handler) Master() Masters {
	return h.master
}

func (h *Handler) Login() Logins {
	return h.login
}
