package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	  _ "github.com/jinzhu/gorm/dialects/sqlite"
	
)

// Models/Struct (Person)
type Person struct {
	gorm.Model
	ID			string `json:"id"`
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
	Birthday	string `json:"birthday"`
	Bio			string `json:"bio"`
}

// Person slice
var persons []Person


// List all person
func listPersons(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

// Get a person
func getPerson(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=  mux.Vars(req) // Get params 

	//Loop and look for the ID
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
		return
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Create a person
func createPerson(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = strconv.Itoa(rand.Intn(10000000)) // @todo change to uuid
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

// Edit a person
func updatePerson(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=  mux.Vars(req)

	for index, item := range persons {
		if item.ID == params["id"]{
		persons = append(persons[:index], persons[index+1:]...)
		var person Person
		_ = json.NewDecoder(req.Body).Decode(&person)
		person.ID = params["id"]
		persons = append(persons, person)
		json.NewEncoder(w).Encode(person)
		return
		}
	}
	json.NewEncoder(w).Encode(persons)

}

// Delete a person
func deletePerson(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range persons {
		if item.ID == params["id"]{
		persons = append(persons[:index], persons[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(persons)

}


func main () {
	//Gorm db settings
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Person{})


	// Router Initialize
	router := mux.NewRouter()
	
	//Data @todo DB
	persons = append(persons, Person{ID: "1", Firstname: "John", Lastname: "Doe", Birthday: "1992-05-26", Bio: "Tall Dark and Handsome"})
	persons = append(persons, Person{ID: "2", Firstname: "Justin", Lastname: "Beiber", Birthday: "1992-06-26", Bio: "Gay"})
	persons = append(persons, Person{ID: "3", Firstname: "Benny", Lastname: "Lee", Birthday: "1996-02-23", Bio: "Shiela"})
	persons = append(persons, Person{ID: "4", Firstname: "Felix", Lastname: "Decutt", Birthday: "1981-02-12", Bio: "Old Soul"})

	// Endpoints
	router.HandleFunc("/api/persons", listPersons).Methods("GET")
	router.HandleFunc("/api/persons/{id}", getPerson).Methods("GET")
	router.HandleFunc("/api/persons", createPerson).Methods("POST")
	router.HandleFunc("/api/persons/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/api/persons/{id}", deletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}