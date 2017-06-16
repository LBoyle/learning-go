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

  // this is where it differs from 02 using gojis
  mux := goji.NewMux()
  mux.HandleFunc(pat.Get("/books"), BooksIndex(session))
  mux.HandleFunc(pat.Post("/books"), BooksCreate(session))
  mux.HandleFunc(pat.Get("/books/:isbn"), BooksByISBN(session))
  mux.HandleFunc(pat.Delete("/books/:isbn"), BooksDelete(session))
  http.ListenAndServe("localhost:4000", mux)
}

func ensureIndex(s *mgo.Session) {
  session := s.Copy()
  defer session.Close()

  // connection, db, collection
  c := session.DB("testgoji").C("books")

  // don't know what background and sparse do db behaviour that you don't have to worry about with JS
  index := mgo.Index{
    Key: []string{"isbn"},
    Unique: true,
    DropDups: true,
    Background: true,
    Sparse: true,
  }
  err := c.EnsureIndex(index)
  if err != nil {
    panic(err)
  }
}

func BooksIndex(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    session := s.Copy()
    defer session.Close()

    c := session.DB("testgoji").C("books")

    var books []Book
    // first find, empty {} for all
    // & what the response will be called
    err := c.Find(bson.M{}).All(&books)
    // bigger error handler
    if err != nil {
      // writer, content, status
      ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
      log.Println("Failed to get all books: ", err)
      // explicit return
      return
    }
    // not sure what MarshalIndent means must google
    respBody, err := json.MarshalIndent(books, "", " ")
    if err != nil {
      log.Fatal(err)
    }
    // success
    ResponseWithJson(w, respBody, httpStatusOK)
  }
}

func BooksCreate(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    session := s.Copy()
    defer session.Close()ar B

    var book Book
    decoder := json.NewDecoder(r.Body)
    if err != nil {
      ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
      return
    }

    c := session.DB("testgoji").C("books")

    err = c.Insert(book)
    if err != nil {
      if mgo.IsDup(err) {
        ErrorWithJSON(w, "Book with this ISBN already exists", http.StatusBadRequest)
        return
      }
      ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
      log.Println("Failed to instert book: ", err)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Location", r.URL.Path+"/"+book.ISBN)
    w.WriteHeader(http.StatusCode)
  }
}

func BooksByISBN(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    session := s.Copy()
    defer session.Close()

    c := session.DB("testgoji").C("books")

  }
}

func BooksUpdate(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    session := s.Copy()
    defer session.Close()

    c := session.DB("testgoji").C("books")

  }
}

func BooksDelete(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    session := s.Copy()
    defer session.Close()

    c := session.DB("testgoji").C("books")

  }
}
