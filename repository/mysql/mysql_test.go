package main

import (
	"database/sql"
	"log"
	"strconv"
	"testing"
	"time"

	r "github.com/moemoe89/go-unit-test-sql/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &r.UserModel{
	ID:          uuid.New().String(),
	FirstName:   "Momo",
	LastName:    "Mock",
	Age:         35,
	DateJoined:  time.Now(),
	DateUpdated: time.Now(),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock){
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection",err)

	}
	return db, mock
}

func TestFindByID(t *testing.T){
	db,mock := NewMock()
	repo := &repository{db}
	defer func(){
		repo.Close()
	}()

	query := "SELECT id, first_name,last_name, age,date_joined,date_updated FROM person  WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id","first_name","last_name", "age","date_joined","date_updated"}).
		AddRow(u.ID, u.FirstName, u.LastName, u.Age, u.DateJoined, u.DateUpdated)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	uId, _ := strconv.Atoi(u.ID)
	user, err := repo.FindByID(uId)
	assert.NotNil(t, user)
	assert.NoError(t,err)
}

func TestFindByIDError(t *testing.T){
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, first_name,last_name, age,date_joined,date_updated FROM person WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id","first_name","last_name", "age","date_joined","date_updated"})

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	uId, _ := strconv.Atoi(u.ID)
	user, err := repo.FindByID(uId)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestFind(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, first_name,last_name, age,date_joined,date_updated FROM person"

	rows := sqlmock.NewRows([]string{"id","first_name","last_name", "age","date_joined","date_updated"}).
		AddRow(u.ID, u.FirstName, u.LastName, u.Age, u.DateJoined, u.DateUpdated)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.Find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO person \\(id, first_name,last_name, age,date_joined,date_updated\\) VALUES \\(\\?, \\?,\\?, \\?,\\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID, u.FirstName, u.LastName, u.Age, u.DateJoined, u.DateUpdated).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(u)
	assert.NoError(t, err)
}

