package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

type tomlConfig struct {
	AmiID          string
	LocalHostname  string
	ProductCodes   string
	ReservationID  string
	PublicHostname string
	PublicIPV4     string
}

var m map[string]string

func main() {

	var config tomlConfig
	if _, err := toml.DecodeFile("metadata.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	m = make(map[string]string)
	m["ami-id"] = config.AmiID
	m["local-hostname"] = config.LocalHostname
	m["reservation-id"] = config.ReservationID
	m["product-codes"] = config.ProductCodes
	m["public-hostname"] = config.PublicHostname
	m["public-ipv4"] = config.PublicIPV4

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
