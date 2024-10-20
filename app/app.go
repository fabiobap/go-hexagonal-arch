package app

import (
	"log"
	"net/http"

	"github.com/go-hexagonal-arch/domain"
	"github.com/go-hexagonal-arch/service"
	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()

	dbClient := getDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDB)}

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
