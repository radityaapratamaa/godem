package user

import "godem/usecase/user"

type Handler struct {
	master Masters
}

func NewHandler(masterUc user.Masters) *Handler {
	return &Handler{master: NewMaster(masterUc)}
}

func (h *Handler) Master() Masters {
	return h.master
}
