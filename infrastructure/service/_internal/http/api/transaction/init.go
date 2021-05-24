package transaction

import "bcg-test/usecase/transaction"

type Handlers interface {
	purchase
}

type Handler struct {
	transactionUc transaction.Usecases
}

func New(transactionUc transaction.Usecases) *Handler {
	return &Handler{
		transactionUc: transactionUc,
	}
}
