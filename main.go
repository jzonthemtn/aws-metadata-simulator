package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var m map[string]string 

func main() {

	m = make(map[string]string)
	m["ami-id"] = "ami-123456789"
	m["local-hostname"] = "ip-10-251-50-12.ec2.internal"
	m["reservation-id"] = "r-fea54097"
	m["product-codes"] = "asdf\n1234"
	m["public-hostname"] = "ec2-203-0-113-25.compute-1.amazonaws.com"
	m["public-ipv4"] = "10.251.50.12"

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/latest/meta-data/{category}", handle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func handle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	i, ok := m[vars["category"]]

	if ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, i)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
