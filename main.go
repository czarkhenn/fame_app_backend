package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	// "math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"time"
	// "github.com/rs/cors"
)

type Response struct {
    Message string `json:"message"`
}


// Models/Struct (Person)
type Person struct {
	//gorm.Model
	ID    	int 		`gorm:"primary_key;auto_increment" json:"id"`
	Firstname	string `gorm:"size:255" json:"firstname"`
	Lastname	string `gorm:"size:255" json:"lastname"`
	Birthday	string `gorm:"size:255" json:"birthday"`
	Bio			string `gorm:"size:255"json:"bio"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}



var db *gorm.DB
var err error

func initDB() {
    db_name := "persons.db"
    db, err = gorm.Open("sqlite3", db_name)

    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }

	// db.Exec("CREATE DATABASE persons")
	// db.Exec("USE persons")
	
    // Migration to create tables for Order and Item schema
    db.AutoMigrate(&Person{})
}


// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
 
// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}



// List all person
func listPersons(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()

	var persons []Person
	db.Find(&persons)
	fmt.Println("{}", persons)

	respondJSON(w, http.StatusOK, persons)
}

// Get a person
func getPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}

	vars := mux.Vars(r)
	key := vars["id"]
	persons := []Person{}
	db.Find(&persons)
	for _, person := range persons {
		// string to int
		s , err:= strconv.Atoi(key)
		if err == nil{
		   if person.ID == s {
			  fmt.Println(person)
			  fmt.Println("Endpoint Hit: Person No:",key)
			  respondJSON(w, http.StatusOK, person)
		   }
		}
	
	}
	 
}

// Create a person
func createPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()
	
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
	}
	keyVal := make(map[string]string)
    json.Unmarshal(body, &keyVal)
    _firstname := keyVal["firstname"]
    _lastname := keyVal["lastname"]
    _birthday := keyVal["birthday"]
	_bio := keyVal["bio"]

	person := &Person{Firstname: _firstname, Lastname: _lastname, Birthday: _birthday, Bio: _bio}
	db.Save(&person)
	respondJSON(w, http.StatusOK, person)
}

// Edit a person
func updatePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	person_id := vars["id"]
	persons := []Person{}
	db.Find(&persons)
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
	}
	s , err := strconv.Atoi(person_id)
	for _, person := range persons {
		// string to int
		if err == nil{
		   if person.ID == s {
			keyVal := make(map[string]string)
			json.Unmarshal(body, &keyVal)
			_firstname := keyVal["firstname"]
			_lastname := keyVal["lastname"]
			_birthday := keyVal["birthday"]
			_bio := keyVal["bio"]
			
			if err != nil {
				resp := Response{Message: "Error please contact admin"}
				respondJSON(w, http.StatusOK, resp)
			}

			fmt.Println(s)
			_person := Person{ID: person.ID, Firstname: _firstname, Lastname: _lastname, Birthday: _birthday, Bio: _bio}
			db.Save(&_person)
			respondJSON(w, http.StatusOK, _person)
		   }

		   
		}
	
	}

}

// Delete a person
func deletePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	key := vars["id"]
	persons := []Person{}
	db.Find(&persons)
	for _, person := range persons {
		// string to int
		s , err := strconv.Atoi(key)
		if err == nil{
		   if person.ID == s {
			  fmt.Println(person)
			  fmt.Println("Endpoint Hit: Person No:",key)
			  resp := Response{Message:"Successfully deleted " + person.Firstname}
			  respondJSON(w, http.StatusOK, resp)
			  db.Delete(&person)
		   }
		
		}
	
	}

}

func deleteAllPerson(w http.ResponseWriter, r *http.Request){
	db, err := gorm.Open("sqlite3", "persons.db")
	if err != nil {
	  panic("failed to connect database")
	}
	defer db.Close()

	var _persons []Person
	db.Find(&_persons)
	db.Delete(&_persons)
	// var deleted = "Deleted Successfully"
	// respondJSON(w, http.StatusOK, deleted)
}







func main () {
	initDB()
	// Router Initialize
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/api/persons", listPersons).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/persons/{id}", getPerson).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/persons", createPerson).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/persons/{id}", updatePerson).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/persons/{id}", deletePerson).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/delete", deleteAllPerson).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8000", r))
	
	
}