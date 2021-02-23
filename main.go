package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// == Models ==
type Person struct {
	//auto increment id
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	// add timestamp
	DateJoined  string `json:"dateJoined"`
	DateUpdated string `json:"dateUpdated"`
	Job         *Job   `json:"job"`
}

type Job struct {
	//auto increment id
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Salary      int    `json:"salary"`
}

// initialize jobs and people as slice
var people []Person

//get all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

//get a single person
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params
	//Loop through people and find with ID, range is used to loop through data structures
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//create a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	// Create random ID - Mock ID
	person.ID = strconv.Itoa(rand.Intn(10000))
	people = append(people, person)
	json.NewEncoder(w).Encode(&Person{})
}

//edit an existing person
func editPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			var person Person
			_ = json.NewDecoder(r.Body).Decode(&person)
			// Create random ID - Mock ID
			person.ID = params["id"]
			people = append(people, person)
			json.NewEncoder(w).Encode(&Person{})
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

//delete a person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

//get all jobs
func getJobs(w http.ResponseWriter, r *http.Request) {

}

//get a single job
func getJob(w http.ResponseWriter, r *http.Request) {

}

//create a new job
func createJob(w http.ResponseWriter, r *http.Request) {

}

//edit an existing job
func editJob(w http.ResponseWriter, r *http.Request) {

}

//delete a job
func deleteJob(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Initialize Router
	r := mux.NewRouter()
	db, err := sql.Open("mysql","test:passw0rd@tcp(localhost:3306)/go_db")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Succesfully connected to MySQL database")

	//jobs as slices @todo implement DB
	people = append(people, Person{ID: "1", FirstName: "Bianca", LastName: "Reusch", Age: 27, DateJoined: "now", DateUpdated: "later", Job: &Job{ID: "1", Title: "Back end Developer", Description: "working on back end", Salary: 45000}})
	people = append(people, Person{ID: "1", FirstName: "Lisa", LastName: "Smith", Age: 29, DateJoined: "now", DateUpdated: "later", Job: &Job{ID: "2", Title: "Front end Developer", Description: "working on front end", Salary: 45000}})

	//Route Handlers / Endpoints
	r.HandleFunc("/people", getPeople).Methods("GET")
	r.HandleFunc("/person/{id}", getPerson).Methods("GET")
	r.HandleFunc("/createPerson", createPerson).Methods("POST")
	r.HandleFunc("/editPerson/{id}", editPerson).Methods("PUT")
	r.HandleFunc("/deletePerson/{id}", deletePerson).Methods("DELETE")

	r.HandleFunc("/jobs", getJobs).Methods("GET")
	r.HandleFunc("/job/{id}", getJob).Methods("GET")
	r.HandleFunc("/createJob", createJob).Methods("POST")
	r.HandleFunc("/editJob/{id}", editJob).Methods("PUT")
	r.HandleFunc("/deleteJob{id}", deleteJob).Methods("DELETE")

	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
