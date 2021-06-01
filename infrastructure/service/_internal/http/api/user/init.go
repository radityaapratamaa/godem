package user

import "godem/usecase/user"

type Handler struct {
	master Masters
	login  Logins
}

func NewHandler(masterUc user.Masters, loginUc user.Logins) *Handler {
	return &Handler{master: NewMaster(masterUc), login: NewLogin(loginUc)}
}

func (h *Handler) Master() Masters {
	return h.master
}

func (h *Handler) Login() Logins {
	return h.login
}
