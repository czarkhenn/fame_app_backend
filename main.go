package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	  _ "github.com/jinzhu/gorm/dialects/sqlite"
	
)




// Models/Struct (Person)
type Person struct {
	gorm.Model
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
	Birthday	string `json:"birthday"`
	Bio			string `json:"bio"`
}

var db *gorm.DB

func initDB() {
    var err error
    db_name := "persons.db"
    db, err = gorm.Open("sqlite3", db_name)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }

    // Migration to create tables for Order and Item schema
    db.AutoMigrate(&Person{})
}

var persons []Person

// List all person
func listPersons(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	
	var persons []Person
	db.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

// Get a person
func getPerson(w http.ResponseWriter, req *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	// params :=  mux.Vars(req) // Get params 

	// //Loop and look for the ID
	// for _, item := range persons {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 	}
	// 	return
	// }
	// json.NewEncoder(w).Encode(&Person{})
}

// Create a person
func createPerson(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	params := mux.Vars(req)
	firstname := params["firstname"]
	lastname := params["lastname"]
	birthday := params["birthday"]
	bio := params["bio"]

	db.Create(&Person{Firstname: firstname, Lastname: lastname, Birthday: birthday, Bio: bio})
	json.NewEncoder(w).Encode(person)
}

// Edit a person
func updatePerson(w http.ResponseWriter, req *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	// params :=  mux.Vars(req)

	// for index, item := range persons {
	// 	if item.ID == params["id"]{
	// 	persons = append(persons[:index], persons[index+1:]...)
	// 	var person Person
	// 	_ = json.NewDecoder(req.Body).Decode(&person)
	// 	person.ID = params["id"]
	// 	persons = append(persons, person)
	// 	json.NewEncoder(w).Encode(person)
	// 	return
	// 	}
	// }
	// json.NewEncoder(w).Encode(persons)

}

// Delete a person
func deletePerson(w http.ResponseWriter, req *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(req)
	// for index, item := range persons {
	// 	if item.ID == params["id"]{
	// 	persons = append(persons[:index], persons[index+1:]...)
	// 	break
	// 	}
	// }
	// json.NewEncoder(w).Encode(persons)

}


func main () {
	// Router Initialize
	router := mux.NewRouter()
	
	// Endpoints
	router.HandleFunc("/api/persons", listPersons).Methods("GET")
	router.HandleFunc("/api/persons/{id}", getPerson).Methods("GET")
	router.HandleFunc("/api/persons", createPerson).Methods("POST")
	router.HandleFunc("/api/persons/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/api/persons/{id}", deletePerson).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
	
	initDB()
}