package app

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/great", great)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
