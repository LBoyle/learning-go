package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "goji.io"
  "goji.io/pat"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

// hello/{name} thing to test goji

// func hello(w http.ResponseWriter, r *http.Request) {
// 	name := pat.Param(r, "name")
// 	fmt.Fprintf(w, "Hello, %s!", name)
// }

// func main() {
// 	mux := goji.NewMux()
// 	mux.HandleFunc(pat.Get("/hello/:name"), hello)
//
// 	http.ListenAndServe("localhost:8000", mux)
// }

// error handler I think
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
  w.Header().Set("Content-Type", "application/json;charset=utf-8")
  w.WriteHeader(code)
  fmt.Fprintf(w, "{message: %q}", message)
}

// response handler
func RespondWithJSON(w http.ResponseWriter, json []byte, code int) {
  w.Header().Set("Content-Type", "application/json;charset=utf-8")
  w.WriteHeader(code)
  fmt.Write(json)
}

// going with books like the tutorial
type Book struct {
  ISBN string `json:"isbn"`
  Title string `json:"title"`
  Authors string `json:"authors"`
  Price string `json:"price"`
}

func main() {
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // dropping here if I need to
  // err = session.DB("testgomongo").DropDatabase()
	// if err != nil {
	// 	panic(err)
	// }

  // saw this in the last tutorial, everybody's doing it so shall I
  // I think it controlls db server traffic
  session.SetMode(mgo.Monotonic, true)
  ensureIndex(session)


}

func ensureIndex(s *mgo.Session) {

}

func BooksIndex(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {

  }
}

func BooksCreate(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {

  }
}

func BooksByISBN(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {

  }
}

func BooksUpdate(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {

  }
}

func BooksDelete(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {

  }
}
