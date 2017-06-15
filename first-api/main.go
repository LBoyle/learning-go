package main

import (
  "encoding/json"
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
  Borough string `json:"borough, omitempty"`
}

var people []Person

func GetPerson(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

func GetPeople(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(people)
}

func CreatePerson(w http.ResponseWriter, req *http.Request) {

}

func DeletePerson(w http.ResponseWriter, req *http.Request) {

}

func main() {
  router := mux.NewRouter()

  people = append(people, Person{ID: "1", Firstname: "Louis", Lastname: "Boyle", Address: &Address{City: "London", Borough: "Islington"}})
  people = append(people, Person{ID: "2", Firstname: "Charlie", Lastname: "Boyle", Address: &Address{City: "London", Borough: "Bow"}})

  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}
