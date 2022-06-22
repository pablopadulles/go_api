package main

import (
	"encoding/json"
	"errors"
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
	router.HandleFunc("/band/members/", GetMember).Methods("GET")
	router.HandleFunc("/band/members/", SetMember).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetBand(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bandsDb)

}

func GetMember(w http.ResponseWriter, r *http.Request) {

	var response []st.Member
	name := r.FormValue("nombre")
	band, err := get_band(bandsDb, name)
	if err != nil {
		response = band.Members
	}
	json.NewEncoder(w).Encode(response)

}

func SetMember(w http.ResponseWriter, r *http.Request) {

	response := false
	nameBand := r.FormValue("nombreBand")
	name := r.FormValue("name")
	lastname := r.FormValue("lastname")

	band, err := get_band(bandsDb, nameBand)
	if err != nil {
		band.Members = append(band.Members, st.Member{Name: name, LastName: lastname})
		response = true
	}
	json.NewEncoder(w).Encode(response)

}

func SetBand(w http.ResponseWriter, r *http.Request) {

	response := false
	name := r.FormValue("nombre")
	if !is_in(name) && name != "" {
		bandsDb = append(bandsDb, st.Band{Name: name})
		response = true
	}
	json.NewEncoder(w).Encode(response)

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

func get_band(bands []st.Band, name string) (st.Band, error) {
	for _, band := range bands {
		if band.Name == name {
			return band, nil
		}
	}
	return st.Band{}, errors.New("La banda no existe")
}
