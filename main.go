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
	ID          int       `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Age         int       `db:"age"`
	DateJoined  time.Time `db:"date_joined"`
	DateUpdated time.Time `db:"date_updated"`
}

type Job struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Salary      int    `db:"salary"`
	EmployeeID 	int	`db:"FK_person"`
}

// == Make connection to DB ==
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "test"
	dbPass := "passw0rd"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/go_db?parseTime=true")
	ErrorCheck(err)
	fmt.Println("Successfully connected to MySQL database")
	return db
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// == Connect to templates ==
var tmpl = template.Must(template.ParseGlob("templates/*"))

//Get all people
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
		var DateUpdated time.Time
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age, &DateJoined, &DateUpdated)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		per.DateJoined = DateJoined
		per.DateUpdated = DateUpdated
		res = append(res, per)
		fmt.Println(res)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//Get a single Person by ID
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
		var DateUpdated time.Time
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age, &DateJoined, &DateUpdated)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		per.DateJoined = DateJoined
		per.DateUpdated = DateUpdated
		fmt.Println(FirstName)
	}
	tmpl.ExecuteTemplate(w, "ShowPerson", per)
	defer db.Close()
}

//Show createPerson Form
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//Insert a new Person to DB
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
		fmt.Println("successfully added person")
	}
	http.Redirect(w, r, "/", 301)
}

//Show EditPerson Form
func showEditPerson(w http.ResponseWriter, r *http.Request) {
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
		var DateUpdated time.Time
		err = selDB.Scan(&ID, &FirstName, &LastName, &Age, &DateJoined, &DateUpdated)
		ErrorCheck(err)
		per.ID = ID
		per.FirstName = FirstName
		per.LastName = LastName
		per.Age = Age
		per.DateJoined = DateJoined
		per.DateUpdated = time.Now()
	}
	tmpl.ExecuteTemplate(w, "Edit", per)
	defer db.Close()
}

//Update a Person in db
func updatePerson(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		FirstName := r.FormValue("firstName")
		LastName := r.FormValue("lastName")
		Age := r.FormValue("age")
		DateUpdated := time.Now()
		ID := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE person SET first_name=?,last_name=?, age=?, date_updated=? WHERE id=?")
		ErrorCheck(err)
		insForm.Exec(FirstName, LastName, Age, DateUpdated, ID)

		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}

//Delete a Person
func deletePerson(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	per := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM person WHERE id=?")
	ErrorCheck(err)
	delForm.Exec(per)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//Get All Jobs
func getJobs(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM job")
	ErrorCheck(err)
	job := Job{}
	var res []Job
	for selDB.Next() {
		var ID int
		var Title, Description string
		var Salary int
		var EmployeeID int
		err = selDB.Scan(&ID, &Title, &Description, &Salary, &EmployeeID)
		ErrorCheck(err)
		job.ID = ID
		job.Title = Title
		job.Description = Description
		job.Salary = Salary
		job.EmployeeID = EmployeeID
		res = append(res, job)

	}
	tmpl.ExecuteTemplate(w, "Jobs", res)
	defer db.Close()
}

//Get a single Job by ID
func getJob(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM job WHERE id=?", nId)
	ErrorCheck(err)
	job := Job{}
	for selDB.Next() {
		var ID int
		var Title, Description string
		var Salary int
		var EmployeeID int
		err = selDB.Scan(&ID, &Title, &Description, &Salary, &EmployeeID)
		ErrorCheck(err)
		job.ID = ID
		job.Title = Title
		job.Description = Description
		job.Salary = Salary
		job.EmployeeID = EmployeeID
	}
	tmpl.ExecuteTemplate(w, "ShowJob", job)
	defer db.Close()
}

//Show Job Form
func NewJob(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewJob", nil)
}

//Insert a new Job into DB
func InsertJob(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Title := r.FormValue("title")
		Description := r.FormValue("description")
		Salary := r.FormValue("salary")
		EmployeeID := 0
		insForm, err := db.Prepare("INSERT INTO job(title, description, salary, FK_person) VALUES(?,?,?,?)")
		ErrorCheck(err)
		insForm.Exec(Title, Description, Salary, EmployeeID)
		defer db.Close()
		fmt.Println("successfully added job")
	}
	http.Redirect(w, r, "/jobs", 301)
}

//Show editJob Form
func showEditJob(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM job WHERE id=?", nId)
	ErrorCheck(err)
	job := Job{}
	for selDB.Next() {
		var ID int
		var Title string
		var Description string
		var Salary int
		var EmployeeID int
		err = selDB.Scan(&ID, &Title, &Description, &Salary, &EmployeeID)
		ErrorCheck(err)
		job.ID = ID
		job.Title = Title
		job.Description = Description
		job.Salary = Salary
		job.EmployeeID = EmployeeID
	}
	tmpl.ExecuteTemplate(w, "EditJob", job)
}

//Edit an existing job
func updateJob(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Title := r.FormValue("title")
		Description := r.FormValue("description")
		Salary := r.FormValue("salary")
		ID := r.FormValue("uid")
		EmployeeID := r.FormValue("employeeID")
		insForm, err := db.Prepare("UPDATE job SET title=?,description=?,salary=?,FK_person=? WHERE id=?")
		ErrorCheck(err)
		insForm.Exec(Title, Description, Salary, ID, EmployeeID)

		defer db.Close()
		http.Redirect(w, r, "/jobs", 301)
	}
}

//Delete a job method
func deleteJob(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	job := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM job WHERE id=?")
	ErrorCheck(err)
	delForm.Exec(job)
	defer db.Close()
	http.Redirect(w, r, "/jobs", 301)
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

	http.HandleFunc("/jobs", getJobs)
	http.HandleFunc("/showJob", getJob)
	http.HandleFunc("/newJob", NewJob)
	http.HandleFunc("/insertJob", InsertJob)
	http.HandleFunc("/editJob", showEditJob)
	http.HandleFunc("/updateJob", updateJob)
	http.HandleFunc("/deleteJob", deleteJob)

	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", r))
}
