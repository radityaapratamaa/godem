package main

import (
	"bcg-test/domain/models"
	"bcg-test/infrastructure/database/transaction"
	"bcg-test/infrastructure/service/_internal/http/api"
	transactionhandler "bcg-test/infrastructure/service/_internal/http/api/transaction"
	"bcg-test/lib/util/database"
	transactionuc "bcg-test/usecase/transaction"
	"log"
)

func loadModules(cfg *models.Config) *api.ModuleHandler {
	mainDB, err := database.Connect(cfg)
	if err != nil {
		log.Fatalln("Cannot connect to DB", err.Error())
	}

	transactionRepo := transaction.New(mainDB)
	transactionUsecase := transactionuc.New(transactionRepo, nil)
	transactionHandler := transactionhandler.New(transactionUsecase)

	return &api.ModuleHandler{
		Transaction: transactionHandler,
	}
}
