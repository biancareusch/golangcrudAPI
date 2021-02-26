package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	repo "github.com/moemoe89/go-unit-test-sql/repository"
	"time"
)

//repository represents the repository model
type repository struct {
	db *sql.DB
}

//NewRepository creates a variable that represents the Repository struct
func NewRepository(dialect, dsn string, idleConn, maxConn int) (repo.Repository, error){
	db, err := sql.Open(dialect,dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

//Close attaches the provider and closes the connection
func (r *repository) Close(){
	r.db.Close()
}

//FindByID attaches the user repository and finds data based on id
func (r *repository) FindByID(id int) (*repo.UserModel, error) {
	user := new(repo.UserModel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}