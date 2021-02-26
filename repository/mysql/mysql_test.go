package main

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	r "github.com/moemoe89/go-unit-test-sql/repository"
	"github.com/stretchr/testify/assert"
)

var u = &r.UserModel{
	IDdb:          uuid.New().String(),
	FirstNamedb:   "Momo",
	LastNamedb:    "Mock",
	Agedb:         35,
	DateJoineddb:  time.Now(),
	DateUpdateddb: time.Now(),
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
		AddRow(u.IDdb, u.FirstNamedb, u.LastNamedb, u.Agedb, u.DateJoineddb, u.DateUpdateddb)

	mock.ExpectQuery(query).WithArgs(u.IDdb).WillReturnRows(rows)

	user, err := repo.FindByID(u.IDdb)
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

	mock.ExpectQuery(query).WithArgs(u.IDdb).WillReturnRows(rows)

	user, err := repo.FindByID(u.IDdb)
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
		AddRow(u.IDdb, u.FirstNamedb, u.LastNamedb, u.Agedb, u.DateJoineddb, u.DateUpdateddb)

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
	prep.ExpectExec().WithArgs(u.IDdb, u.FirstNamedb, u.LastNamedb, u.Agedb, u.DateJoineddb, u.DateUpdateddb).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(u)
	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO person \\(id, first_name,last_name, age,date_joined,date_updated\\) VALUES \\(\\?, \\?, \\?,\\?,\\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.IDdb, u.FirstNamedb, u.LastNamedb, u.Agedb, u.DateJoineddb, u.DateUpdateddb).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Create(u)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE person SET name = \\?, first_name = \\?,last_name = \\?,age = \\?, date_joined = \\?,date_updated = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.IDdb, u.FirstNamedb, u.LastNamedb, u.Agedb, u.DateJoineddb, u.DateUpdateddb).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(u)
	assert.NoError(t, err)
}

func TestUpdateErr(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE person SET name = \\?, first_name = \\?,last_name = \\?,age = \\?, date_joined = \\?,date_updated = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.IDdb, u.FirstNamedb, u.LastNamedb, u.Agedb, u.DateJoineddb, u.DateUpdateddb).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Update(u)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM person WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.IDdb).WillReturnResult(sqlmock.NewResult(0, 1))


	err := repo.Delete(u.IDdb)
	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM person WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.IDdb).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(u.IDdb)
	assert.Error(t, err)
}