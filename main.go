package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
);

func main()  {
    router := mux.NewRouter()

    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}

type Person struct {
    ID        string `json:"id,omitempty"`
    FirstName string `json:"firstName,omitempty"`
    LastName  string `json:"lastName,omitempty"`
    Address  *Address`json:"address,omitempty"`
}

type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
        }
    }
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index + 1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(people)
}
