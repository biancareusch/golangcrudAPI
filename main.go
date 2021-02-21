package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// == Models ==
type Person struct {
	//auto increment id
	ID        int    `json:"id"`
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
	ID          int    `json:"id"`
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

}

//create a new person
func createPerson(w http.ResponseWriter, r *http.Request) {

}

//edit an existing person
func editPerson(w http.ResponseWriter, r *http.Request) {

}

//delete a person
func deletePerson(w http.ResponseWriter, r *http.Request) {

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

	//jobs as slices @todo implement DB
	people = append(people, Person{ID: 1, FirstName: "Bianca", LastName: "Reusch", Age: 27, DateJoined: "now", DateUpdated: "later", Job: &Job{ID: 1, Title: "Back end Developer", Description: "working on back end", Salary: 45000}})
	people = append(people, Person{ID: 2, FirstName: "Lisa", LastName: "Smith", Age: 29, DateJoined: "now", DateUpdated: "later", Job: &Job{ID: 2, Title: "Front end Developer", Description: "working on front end", Salary: 45000}})

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