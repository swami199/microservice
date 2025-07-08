package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func great(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Nuvu super ra bulloda")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "swami babu", City: "Parlakhemundi", Zipcode: "761200"},
		{Name: "varun", City: "Parlakhemundi", Zipcode: "761200"},
	}
	if r.Header.Get("Content-Type") == "Application/json" {
		w.Header().Add("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "Application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}
func getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post")
}
