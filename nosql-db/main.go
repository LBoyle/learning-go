package main


import (
  "fmt"
	"log"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
  Name string
  Email string
  Address *Address
}

type Address struct {
  City string
  Borough string
}

func main() {
  // connect to database
  session, err := mgo.Dial("mongodb://localhost:27017")
  if err != nil {
    panic(err)
  }
  // don't know what this means
  defer session.Close()
  // c for connection? is my guess
  c := session.DB("testgomongo").C("people")
  err = c.Insert(&Person{"Louis", "louis@louis.com", {"London", "Islington"}}, &Person{"Gabe", "gabe@gabe.com", {"London", "Islington"}})
  if err != nil {
    log.Fatal(err)
  }
  // getting a person out by name
  result := Person{}
  err = c.Find(bson.M{"name": "Louis"}).One(&result)
  if err != nil {
    log.Fatal(err)
  }
  // prove that it works
  fmt.Println("Email: ", result.Email)
}
