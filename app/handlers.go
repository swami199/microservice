package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

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
