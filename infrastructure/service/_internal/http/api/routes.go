package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bcg-test/infrastructure/service/_internal/http/api/transaction"
)

type Routes struct {
	router  *mux.Router
	modules *ModuleHandler
}

type ModuleHandler struct {
	Transaction transaction.Handlers
}

func NewHandler(modules *ModuleHandler) *Routes {
	return &Routes{
		router:  mux.NewRouter().StrictSlash(true),
		modules: modules,
	}
}

func (h *Routes) RegisterAndStartServer() error {

	//register your routes here
	h.router.HandleFunc("/ping", h.Ping)

	// transaction routes start
	h.router.HandleFunc("/transaction/purchase", h.modules.Transaction.CreateNew)

	log.Println("http listening on port :10000")
	return http.ListenAndServe(":10000", h.router)
}

func (h *Routes) Ping(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"code":    200,
		"message": "OK",
	}

	json.NewEncoder(w).Encode(response)
}
