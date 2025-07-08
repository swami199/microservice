package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//define routers
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/great", great).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/createCustomer", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomerById)
	//definimg server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
