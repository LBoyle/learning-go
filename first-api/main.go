package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

type Person struct {
  ID string `json:"id, omitempty"`
  Firstname string `json:"firstname, omitempty"`
  Lastname string `json:"lastname, omitempty"`
  Address *Address `json:"address, omitempty"`
}

type Address struct {
  City string `json:"city, omitempty"`
  Borough string `json:"city, omitempty"`
}

func GetPerson(w http.ResponseWriter, req *http.Request) {

}

func GetPeople(w http.ResponseWriter, req *http.Request) {

}

func CreatePerson(w http.ResponseWriter, req *http.Request) {

}

func DeletePerson(w http.ResponseWriter, req *http.Request) {

}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePeople).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}
