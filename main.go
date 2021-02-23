package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
	_ "text/template"
	"time"
)

// == Models ==
type Person struct {
	ID         int       `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Age        int       `db:"age"`
	DateJoined time.Time `db:"date_joined"`
}

type Job struct {
	//auto increment id
	ID          string
	Title       string
	Description string
	Salary      int
}

//make connection to DB
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "test"
	dbPass := "passw0rd"
	db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/go_db?parseTime=true")
	ErrorCheck(err)
	fmt.Println("Succesfully connected to MySQL database")
	return db
}

func ErrorCheck(err error){
	if err != nil {
		panic(err.Error())
	}
}
// get views
var tmpl = template.Must(template.ParseGlob("templates/*"))

//get all people

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM person")
	ErrorCheck(err)
	per := Person{}
	var res []Person
	for selDB.Next() {
		var ID int
		var FirstName, LastName string
		var Age int
		var DateJoined time.Time
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age, &DateJoined)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		per.DateJoined = DateJoined
		res = append(res, per)
		fmt.Println(res)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//get a single person
func showPerson(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM person WHERE id=?", nId)
	ErrorCheck(err)
	per := Person{}
	for selDB.Next() {
		var ID int
		var FirstName, LastName string
		var Age int
		var DateJoined time.Time
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age, &DateJoined)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		per.DateJoined = DateJoined
		fmt.Println(FirstName)
	}
	tmpl.ExecuteTemplate(w, "ShowPerson", per)
	defer db.Close()
}

//show create new Person Form
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//create a new person
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		FirstName := r.FormValue("firstName")
		LastName := r.FormValue("lastName")
		Age := r.FormValue("age")
		DateJoined := time.Now()
		insForm, err := db.Prepare("INSERT INTO person(first_name, last_name, age,date_joined) VALUES(?,?,?,?)")
		ErrorCheck(err)
		insForm.Exec(FirstName, LastName, Age, DateJoined)
	defer db.Close()
		fmt.Println("succesfully added person")
	}
	http.Redirect(w, r, "/", 301)
}

//show edit
func showEditPerson(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM person WHERE ID=?", nId)
	ErrorCheck(err)
	per := Person{}
	for selDB.Next() {
		var ID int
		var FirstName, LastName string
		var Age int
		var DateJoined time.Time
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age, &DateJoined)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		per.DateJoined = DateJoined
	}
	tmpl.ExecuteTemplate(w, "Edit", per)
	defer db.Close()
}

//update a person
func updatePerson(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		FirstName := r.FormValue("firstName")
		LastName := r.FormValue("lastName")
		Age := r.FormValue("age")
		DateJoined := time.Now()
		ID := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE person SET first_name=?,last_name=?, age=?, date_joined=? WHERE id=?")
		ErrorCheck(err)
		insForm.Exec(FirstName, LastName, Age, DateJoined, ID)

		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}

//delete a person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	per := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM person WHERE id=?")
	ErrorCheck(err)
	delForm.Exec(per)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
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

	//Route Handlers / Endpoints
	http.HandleFunc("/", Index)
	http.HandleFunc("/showPerson", showPerson)
	http.HandleFunc("/newPerson", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/editPerson", showEditPerson)
	http.HandleFunc("/updatePerson", updatePerson)
	http.HandleFunc("/deletePerson", deletePerson)

	r.HandleFunc("/jobs", getJobs).Methods("GET")
	r.HandleFunc("/job/{id}", getJob).Methods("GET")
	r.HandleFunc("/createJob", createJob).Methods("POST")
	r.HandleFunc("/editJob/{id}", editJob).Methods("PUT")
	r.HandleFunc("/deleteJob{id}", deleteJob).Methods("DELETE")

	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", r))
}
