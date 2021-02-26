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

	err := r.db.QueryRowContext(ctx, "SELECT id, first_name, last_name,age,date_joined,date_updated FROM person WHERE id = ?").Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.DateJoined,&user.DateUpdated)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//Find attaches the user repository and finds its data
func (r *repository) Find()([]*repo.UserModel,error){
	users := make([]*repo.UserModel,0)

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id, first_name, last_name,age,date_joined,date_updated FROM person")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		user := new(repo.UserModel)
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Age,
			&user.DateJoined,
			&user.DateUpdated,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

//Create attaches the user repository and creates data
func (r *repository) Create(user *repo.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO person(id, first_name,last_name, age,date_joined,date_updated) VALUES(?,?,?,?,?,?)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_,err = stmt.ExecContext(
		ctx,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Age,
		user.DateJoined,
		user.DateUpdated)
	return err
}

//Delete attaches the user repository and deletes data found by ID
func (r *repository) Delete(id int) error{
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	query := "DELETE FROM person WHERE id = ?"
	stmt, err := r.db.PrepareContext(ctx,query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
