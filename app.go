package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	st "go_api/band_struct"

	"github.com/gorilla/mux"
)

var bandsDb []st.Band

func main() {
	fmt.Println("Empeso todo")
	bandsDb = append(bandsDb, st.Band{Name: "ALMAFUERTE"})

	router := mux.NewRouter()

	router.HandleFunc("/band/", GetBand).Methods("GET")
	router.HandleFunc("/band/", SetBand).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetBand(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bandsDb)

}

func SetBand(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("nombre")
	if !is_in(name) {
		bandsDb = append(bandsDb, st.Band{Name: name})
	}
	json.NewEncoder(w).Encode(nil)

}

func Search(Name string) st.Band {

	for _, band := range bandsDb {
		if band.Name == Name {
			return band
		}
	}
	return st.Band{}
}

func is_in(name string) bool {
	for _, band := range bandsDb {
		if band.Name == name {
			return true
		}
	}
	return false
}
